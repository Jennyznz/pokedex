package main

import ( 
	"fmt"
	"bufio"
	"os"
	"github.com/Jennyznz/pokedex/internal/pokecache"
	"github.com/Jennyznz/pokedex/internal/pokeapi"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	cleaned := []string{}
	command := ""
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(cache)
	c := &config{
		pokeapiClient: client,
		pokedex: make(map[string]pokemon),
	} 

	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() == true {
			input = scanner.Text()
			cleaned = cleanInput(input)
			command = cleaned[0]
			verifiedCommand, exists := commands[command]
			
			if exists == true {
				if len(cleaned) == 1 {
					// Zero parameter callbacks
					verifiedCommand.callback(c)
				} else {
					argOne := cleaned[1]
					verifiedCommand.callback(c, argOne)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

