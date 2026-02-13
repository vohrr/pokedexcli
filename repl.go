package main

import (
	"fmt"
	"github.com/vohrr/pokeapi"
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
	dex      map[string]pokeapi.PokemonResponse
	Next     *string
	Previous *string
}

const (
	exit    string = "exit"
	help    string = "help"
	map_cmd string = "map"
	mapb    string = "mapb"
	explore string = "explore"
	catch   string = "catch"
	inspect string = "inspect"
	pokedex string = "pokedex"
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
		inspect: {
			name:        inspect,
			description: "View data of a captured Pokemon",
			callback:    commandInspect,
		},
		catch: {
			name:        catch,
			description: "Attempt to catch the specified Pokemon",
			callback:    commandCatch,
		},
		explore: {
			name:        explore,
			description: "Retrieve detailed information about a specific location in Pokemon",
			callback:    commandExplore,
		},
		pokedex: {
			name:        pokedex,
			description: "List the Pokemon you have caught so far",
			callback:    commandPokedex,
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

func extractArg(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("location area not specified")
	}
	return args[0], nil
}
