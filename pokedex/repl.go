package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := config{"https://pokeapi.co/api/v2/location-area?offset=0&limit=20", ""}
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Print("\n\n")
	commands := getCommands()

	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationResponse, err := getLocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locationResponse.Next
	if locationResponse.Previous != nil {
		cfg.Previous = *locationResponse.Previous
	}
	return nil
}

func getLocationAreas(url string) (locationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer res.Body.Close()
	locationResponse := locationAreaResponse{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return locationAreaResponse{}, err
	}
	for _, l := range locationResponse.Results {
		fmt.Println(l.Name)
	}
	return locationResponse, nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResponse, err := getLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = cfg.Previous
	if locationResponse.Previous != nil {
		cfg.Previous = *locationResponse.Previous
	} else {
		cfg.Previous = ""
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Next     string
	Previous string
}
