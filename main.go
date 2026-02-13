package main

import (
	"bufio"
	"fmt"
	"github.com/vohrr/pokeapi"
	"github.com/vohrr/pokecache"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var config config
	config.cache = pokecache.NewCache(60 * time.Second)
	config.dex = make(map[string]pokeapi.PokemonResponse)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {

			input := cleanInput(scanner.Text())

			if len(input) < 1 {
				fmt.Println("Invalid Input")
				continue
			}

			var err error
			commands := getCommands()

			if command, ok := commands[input[0]]; ok {
				err = command.callback(&config, input[1:]...)

				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Unknown command")
				commands[help].callback(&config)
			}

		} else {
			fmt.Println("Invalid Input")
		}
	}
}
