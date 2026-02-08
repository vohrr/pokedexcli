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
	callback    func() error
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
			callback:    commandMap,
		},
		exit: {
			name:        exit,
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap() error {
	// call pokeAPI to grab locations
	response, err := pokeapi.Fetch[pokeapi.LocationAreaResponse](pokeapi.GetLocationAreasUrl)
	if err != nil {
		return err
	}
	fmt.Printf("Found %d areas!\n", response.Count)
	return nil
}
