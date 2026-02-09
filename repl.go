package main

import (
	"github.com/vohrr/pokecache"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, arg ...string) error
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
	explore string = "explore"
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
			description: "Retrieves and displays the next page of areas in the Pokemon world",
			callback:    commandMapf,
		},
		mapb: {
			name:        mapb,
			description: "Retrieves and displays the previous page of areas in the Pokemon world",
			callback:    commandMapb,
		},
		explore: {
			name:        explore,
			description: "Retrieve detailed information about a specific location in Pokemon",
			callback:    commandExplore,
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
