package main

import (
	"fmt"
)

func commandInspect(c *config, args ...string) error{
	pokemon, exists := c.pokedex[args[0]]
	if exists {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats: ")
		for _, s := range pokemon.Stats {
			fmt.Printf("     - %s: %d\n", s.Stat.Name, s.BaseVal)
		}
		fmt.Println("Types: ")
		for _, t := range pokemon.Types {
			fmt.Println("     - " + t.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon.")
	}

	return nil
}