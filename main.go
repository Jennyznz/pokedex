package main

import ( 
	"fmt"
	"bufio"
	"os"
)

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