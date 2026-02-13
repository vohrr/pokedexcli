package main

import "fmt"

func commandPokedex(config *config, args ...string) error {

	if len(config.dex) == 0 {
		fmt.Println("You haven't caught any Pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.dex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
