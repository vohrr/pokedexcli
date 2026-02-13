package main

import "fmt"

func commandInspect(config *config, args ...string) error {
	pokemon, err := extractArg(args)
	if err != nil {
		return err
	}

	if data, ok := config.dex[pokemon]; ok {
		//print data about the pokemon
		fmt.Printf("Name: %s\n", data.Name)
		fmt.Printf("Height: %d\n", data.Height)
		fmt.Printf("Weight: %d\n", data.Weight)
		fmt.Println("Stats:")
		for _, stat := range data.Stats {
			fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, poketype := range data.Types {
			fmt.Printf("	- %s\n", poketype.Type.Name)
		}

	} else {
		fmt.Println("You haven't caught that Pokemon yet")
	}

	return nil
}
