package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean := cleanInput(scanner.Text())
		if len(clean) == 0 {
			continue
		}
		command, ok := getCommands()[clean[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&cfg)
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
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}
