package main

import (
	"fmt"
	"io"
	"encoding/json"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/ability/" + args[0] + "/"

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
			return fmt.Errorf("Failed to read from response body")
		}

		defer res.Body.Close()
		c.pokeapiClient.Cache.Add(url, body)
	}

	var thePokemon pokemon
	if err := json.Unmarshal(body, &thePokemon); err != nil {
		return fmt.Errorf("Failed to unload JSON data")
	}

	fmt.Printf("\nThrowing a Pokeball at %s...", args[0])

	const avgBaseExp = 200
	randVal := rand.Intn(thePokemon.BaseExperience)
	if randVal < avgBaseExp {
		fmt.Println("%s was caught!", args[0])
	} else {
		fmt.Println("%s escaped!", args[0])
	}

	return nil
}