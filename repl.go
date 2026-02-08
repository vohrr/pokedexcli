package main

import (
	"fmt"
	"github.com/vohrr/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	Next     *string
	Previous *string
}

const (
	exit    string = "exit"
	help    string = "help"
	map_cmd string = "map"
	mapb    string = "mapb"
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
	//cache check

	// call pokeAPI to grab locations on cache miss
	url := pokeapi.LocationAreasUrl
	if config.Next != nil {
		url = *config.Next
	}
	response, err := pokeapi.Fetch[pokeapi.LocationAreaResponse](url)
	if err != nil {
		return err
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
		fmt.Println("you're on the first page")
		return nil
	}
	response, err := pokeapi.Fetch[pokeapi.LocationAreaResponse](*config.Previous)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	config.Next = response.Next
	config.Previous = response.Previous
	return nil
}
