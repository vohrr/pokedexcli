package main

import (
	"encoding/json"
	"fmt"
	"github.com/vohrr/pokeapi"
)

func commandMapf(config *config, args ...string) error {
	url := pokeapi.LocationAreasUrl
	if config.Next != nil {
		url = *config.Next
	}
	err := getMap(config, url)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(config *config, args ...string) error {
	if config.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	err := getMap(config, *config.Previous)
	if err != nil {
		return err
	}
	return nil
}

func getMap(config *config, url string) error {

	var response pokeapi.LocationAreaResponse
	var err error
	value, ok := config.cache.Get(url)
	if !ok {
		response, err = pokeapi.Fetch[pokeapi.LocationAreaResponse](url)
		if err != nil {
			return err
		}

		value, err = json.Marshal(response)
		err = config.cache.Add(url, value)

		if err != nil {
			return err
		}
	} else {
		err = json.Unmarshal(value, &response)
		if err != nil {
			return err
		}
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	config.Next = response.Next
	config.Previous = response.Previous
	return nil

}
