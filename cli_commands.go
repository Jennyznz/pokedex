package main

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	Next string 
	Previous string
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