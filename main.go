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

	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() == true {
			input = scanner.Text()
			cleaned = cleanInput(input)
			fmt.Println("Your command was:", cleaned[0])
		}
	}
}
