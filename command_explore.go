package main

import (
	"fmt"
	"encoding/json"
	"io"
	"strings"
)

func commandExplore(c *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("The Explore command requires a location input")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + args[0] + "/"
	
	var body []byte
	val, ok := c.pokeapiClient.Cache.Get(url)
	if ok {
		body = val
	} else {
		res, err := c.pokeapiClient.HttpClient.Get(url)
		if err != nil {
			return fmt.Errorf("Failed to get data from PokeAPI")
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Error reading from response body")
		}
		
		defer res.Body.Close()
		c.pokeapiClient.Cache.Add(url, body)
	}

	var data location
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("Failed to unload JSON data")
	}

	fmt.Printf("Exploring %s...\n", strings.ToLower(args[0]))
	fmt.Println("Found Pokemon: ")

	for _, encounter := range data.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}

