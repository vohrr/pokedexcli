package main

import (
	"encoding/json"
	"fmt"
	"github.com/vohrr/pokeapi"
	"math/rand/v2"
)

func commandCatch(config *config, args ...string) error {
	pokemon, err := extractArg(args)
	if err != nil {
		return err
	}

	var res pokeapi.PokemonResponse

	url := fmt.Sprintf("%s/%s", pokeapi.PokemonUrl, pokemon)

	value, ok := config.cache.Get(url)
	if !ok {

		res, err = pokeapi.Fetch[pokeapi.PokemonResponse](url)
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

	fmt.Printf("Throwing a Pokeball at %s....\n", pokemon)

	//calculate random chance to catch base on base xp
	baseChance := 100 / (1 + res.BaseExperience/100)
	if baseChance > rand.IntN(100) {
		fmt.Printf("%s was caught!\n", pokemon)
		//add caught poke to dex so that we can inspect it later
		config.dex[pokemon] = res
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
