package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var config config

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

				err = command.callback(&config)

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
