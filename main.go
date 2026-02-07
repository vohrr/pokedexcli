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
			switch input[0] {
			case exit:
				commands[exit].callback()
			case help:
				commands[help].callback()
			default:
				fmt.Println("Unknown command")
			}
		} else {
			fmt.Printf("Invalid Input")
		}
	}
}
