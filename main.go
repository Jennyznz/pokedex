package main

import ( 
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	cleaned := []string{}
	command := ""

	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() == true {
			input = scanner.Text()
			cleaned = cleanInput(input)
			command = cleaned[0]
			verifiedCommand, exists := commands[command]
			if exists == true {
				verifiedCommand.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
