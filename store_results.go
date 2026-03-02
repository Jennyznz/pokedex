package main

type locationArea struct {
	Next *string	`json:"next"`
	Previous *string	`json:"previous"`
	Results []struct {
		Name string	`json:"name"`
		URL string	`json:"url"`
	}	`json:"results"`
}

type location struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:name`
			URL string	`json:url`
		} `json:pokemon`
	} `json:"pokemon_encounters"`
}

type pokemon struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
}