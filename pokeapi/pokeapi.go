package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	Url  string
}

const (
	GetLocationAreasUrl string = "https://pokeapi.co/api/v2/location-area"
)

func Fetch[T any](url string) (T, error) {
	res, err := http.Get(url)
	var returnStruct T
	if err != nil || res.StatusCode > 299 {
		return returnStruct, fmt.Errorf("Failed request to POKEAPI, %s", res.Status)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&returnStruct)

	if err != nil {
		return returnStruct, err
	}

	return returnStruct, nil
}
