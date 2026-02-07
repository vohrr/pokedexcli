package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := cleanInput(scanner.Text())
			commands := getCommands()

			if command, ok := commands[input[0]]; ok {
				command.callback()
			} else {
				fmt.Println("Unknown command")
				commands[help].callback()
			}
		} else {
			fmt.Printf("Invalid Input")
		}
	}
}
