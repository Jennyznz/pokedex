package main

import ( 
	"fmt"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	Next string 
	Previous string
}

type locationArea struct {
	Next *string	`json:"next"`
	Previous *string	`json:"previous"`
	Results []struct {
		Name string	`json:"name"`
		URL string	`json:"url"`
	}	`json:"results"`
}

var commands = map[string]cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays available commands",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays 20 location areas in the Pokemon world. Consecutive calls display next 20 location areas.",
		callback: commandMap,
	},
	"mapb":  {
		name: "mapb",
		description: "Displays previous 20 location areas",
		callback: commandMapB,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	cleaned := []string{}
	command := ""
	c := &config{} // An empty pointer to a config struct

	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() == true {
			input = scanner.Text()
			cleaned = cleanInput(input)
			command = cleaned[0]
			verifiedCommand, exists := commands[command]
			if exists == true {
				verifiedCommand.callback(c)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Next != "" {
		url = c.Next
	}

	res, err := http.Get(url)
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

	res, err := http.Get(url)
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

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays location areas in the Pokemon world")
	return nil
}
