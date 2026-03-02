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
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		Stat struct {
			Name string `json:"name"`
		}	`json:"stat"`
		BaseVal int `json:"base_stat"`
	}	`json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		}	`json:"type"`
	}	`json:"types"`
}