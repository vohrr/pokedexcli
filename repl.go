package main

import (
	"encoding/json"
	"fmt"
	"github.com/vohrr/pokeapi"
	"github.com/vohrr/pokecache"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	cache    *pokecache.Cache
	Next     *string
	Previous *string
}

const (
	exit    string = "exit"
	help    string = "help"
	map_cmd string = "map"
	mapb    string = "mapb"
)
const (
	logCache bool = false
)

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		help: {
			name:        help,
			description: "Displays a help message",
			callback:    commandHelp,
		},
		map_cmd: {
			name:        "map",
			description: "Lists the areas of the Pokemon world",
			callback:    commandMap,
		},
		mapb: {
			name:        mapb,
			description: "",
			callback:    commandMapb,
		},
		exit: {
			name:        exit,
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	return strings.Fields(strings.ToLower(text))
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(config *config) error {
	url := pokeapi.LocationAreasUrl
	if config.Next != nil {
		url = *config.Next
	}

	//cache check
	var response pokeapi.LocationAreaResponse
	var err error
	value, ok := config.cache.Get(url)
	if !ok {
		if logCache {
			pokecache.CacheLog(false, url)
		}
		// call pokeAPI to grab locations on cache miss
		response, err = pokeapi.Fetch[pokeapi.LocationAreaResponse](url)
		if err != nil {
			return err
		}

		value, err = json.Marshal(response)
		if err != nil {
			return err
		}

		err = config.cache.Add(url, value)
		if err != nil {
			return err
		}

	} else {
		if logCache {
			pokecache.CacheLog(true, url)
		}
		err = json.Unmarshal(value, &response)
		if err != nil {
			return err
		}
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	config.Next = response.Next
	config.Previous = response.Previous
	return nil
}

func commandMapb(config *config) error {
	if config.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	var response pokeapi.LocationAreaResponse
	var err error
	value, ok := config.cache.Get(*config.Previous)
	if !ok {
		// call pokeAPI to grab locations on cache miss
		if logCache {
			pokecache.CacheLog(false, *config.Previous)
		}
		response, err = pokeapi.Fetch[pokeapi.LocationAreaResponse](*config.Previous)
		if err != nil {
			return err
		}

		value, err = json.Marshal(response)
		err = config.cache.Add(*config.Previous, value)

		if err != nil {
			return err
		}
	} else {
		if logCache {
			pokecache.CacheLog(true, *config.Previous)
		}
		err = json.Unmarshal(value, &response)
		if err != nil {
			return err
		}
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	config.Next = response.Next
	config.Previous = response.Previous
	return nil
}
