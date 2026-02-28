package main

import (
	"fmt"
	"encoding/json"
	"io"
)

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Next != "" {
		url = c.Next
	}

	res, err := c.pokeapiClient.HttpClient.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to get data from PokeAPI")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("Failed to read response body")
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


func commandMapB(c *config) error {
	url := ""

	if c.Previous != "" {
		url = c.Previous
	} else {
		return fmt.Errorf("You are on the first page")
	}

	res, err := c.pokeapiClient.HttpClient.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to get data from PokeAPI")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code %d", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("Failed to read response body")
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
