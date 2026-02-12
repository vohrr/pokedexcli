package main

import (
	"encoding/json"
	"fmt"
	"github.com/vohrr/pokeapi"
)

func commandExplore(config *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("location area not specified")
	}
	url := fmt.Sprintf("%s/%s", pokeapi.LocationAreasUrl, args[0])
	var res pokeapi.LocationAreaResponseSingle
	var err error
	value, ok := config.cache.Get(url)
	if !ok {
		res, err = pokeapi.Fetch[pokeapi.LocationAreaResponseSingle](url)
		if err != nil {
			return err
		}

		value, err = json.Marshal(res)
		err = config.cache.Add(url, value)
		if err != nil {
			return err
		}
	} else {
		err = json.Unmarshal(value, &res)
		if err != nil {
			return err
		}
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
