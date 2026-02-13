package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	LocationAreasUrl string = "https://pokeapi.co/api/v2/location-area"
	PokemonUrl       string = "https://pokeapi.co/api/v2/pokemon"
)

func Fetch[T any](url string) (T, error) {
	res, err := http.Get(url)
	var responseObject T
	if err != nil || res.StatusCode > 299 {
		return responseObject, fmt.Errorf("Failed request to POKEAPI, %s", res.Status)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&responseObject)

	if err != nil {
		return responseObject, err
	}

	return responseObject, nil
}
