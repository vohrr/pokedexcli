package main

import (
	"fmt"
	"github.com/vohrr/pokeapi"
)

// "encoding/json"
/// "github.com/vohrr/pokeapi"/ "github.com/vohrr/pokecache"

func commandExplore(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("location area not specified")
	}
	url := fmt.Sprintf("%s/%s", pokeapi.LocationAreasUrl, args[0])

	res, err := pokeapi.Fetch[pokeapi.LocationAreaResponseSingle](url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	if res.PokemonEncounters == nil || len(res.PokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in this area!")
	}

	fmt.Printf("Found Pokemon:\n")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
