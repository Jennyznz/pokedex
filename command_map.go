package main

import (
	"fmt"
	"encoding/json"
	"io"
)

func commandMap(c *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Next != "" {
		url = c.Next
	}

	var body []byte
	// Check cache
	val, ok := c.pokeapiClient.Cache.Get(url) 
	if ok {
		body = val
	} else {
		res, err := c.pokeapiClient.HttpClient.Get(url)
		if err != nil {
			return fmt.Errorf("Failed to get data from PokeAPI")
		}
		body, err = io.ReadAll(res.Body) 
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
		}
		if err != nil {
			return fmt.Errorf("Failed to read response body")
		}
		defer res.Body.Close()	// Is function-scoped
		c.pokeapiClient.Cache.Add(url, body)
	}

	var data locationArea
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("Failed to unload JSON data")
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	if data.Previous != nil {
		c.Previous = *data.Previous
	} else {
		c.Previous = ""
	}

	if data.Next != nil {
		c.Next = *data.Next
	} else {
		c.Next = ""
	}

	return nil
}


func commandMapB(c *config, args ...string) error {
	url := ""

	if c.Previous != "" {
		url = c.Previous
	} else {
		return fmt.Errorf("You are on the first page")
	}

	var body []byte
	// Check cache
	val, ok := c.pokeapiClient.Cache.Get(url)
	if ok {
		body = val
	} else {
		res, err := c.pokeapiClient.HttpClient.Get(url)
		if err != nil {
			return fmt.Errorf("Failed to get data from PokeAPI")
		}

		body, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code %d", res.StatusCode)
		}
		if err != nil {
			return fmt.Errorf("Failed to read response body")
		}
		defer res.Body.Close()
		c.pokeapiClient.Cache.Add(url, body)
	}

	var data locationArea
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("Failed to unload JSON data") 
	}

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	if data.Previous != nil {
		c.Previous = *data.Previous
	} else {
		c.Previous = ""
	}

	if data.Next != nil {
		c.Next = *data.Next
	} else {
		c.Next = ""
	}

	return nil
}
