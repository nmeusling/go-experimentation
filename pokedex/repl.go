package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nmeusling/go-experimentation/pokedex/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean := cleanInput(scanner.Text())
		if len(clean) == 0 {
			continue
		}
		parameters := clean[1:]
		command, ok := getCommands()[clean[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, parameters)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	words := []string{}
	fields := strings.Fields(text)

	for _, field := range fields {
		words = append(words, strings.ToLower(field))
	}
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon present in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon with a Pokeball",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays details of a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all pokemon that you have in your pokedex",
			callback:    commandPokedex,
		},
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	Next          string
	Previous      string
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokemon
}

type pokemon struct {
	name           string
	baseExperience int
	weight         int
	height         int
	stats          map[string]int
	types          []string
}

func (p pokemon) printPokemonData() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Height: %v\n", p.height)
	fmt.Printf("Weight: %v\n", p.weight)
	fmt.Println("Stats:")
	for name, val := range p.stats {
		fmt.Printf("  -%s: %v\n", name, val)
	}
	fmt.Println("Types:")
	for _, pokeType := range p.types {
		fmt.Printf("  - %s\n", pokeType)
	}
}
