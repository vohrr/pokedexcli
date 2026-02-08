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
			var err error
			if command, ok := commands[input[0]]; ok {
				err = command.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
				commands[help].callback()
			}
		} else {
			fmt.Printf("Invalid Input")
		}
	}
}
