# Pokedex

## Overview
This project is a command-line Pokédex built in Go. It allows users to explore Pokémon data directly from the terminal by interacting with an external API. The goal of the project is to practice Go fundamentals such as HTTP requests, JSON parsing, concurrency, and building interactive CLI applications.

## Features
- Browse Pokémon by location areas
- Catch and store Pokémon in your personal Pokédex
- View detailed information about captured Pokémon
- Pagination support for exploring large datasets
- Interactive command-based REPL interface
- Caching system to reduce redundant API calls
- Clean and modular Go code structure for easy extension

## Usage 
Once the application is running, you can interact with it using the following commands:
- map: displays 20 locations
- mapb: map back command. Displays previous 20 lcoations.
- explore <location area>: lists all the Pokemon at specified location
- catch <pokemon name>: attempts to catch specified Pokemon
- inspect <pokemon name>: prints name, height, weight, stats, and type(s) of Pokemon
- pokedex: prints list of all Pokemon caught by user
- help: shows available commands
- exit: exits the program
