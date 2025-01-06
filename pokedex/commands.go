package main

import (
	"fmt"
	"os"

	"github.com/nmeusling/go-experimentation/pokedex/internal/pokeapi"
)

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
	locationResponse, err := pokeapi.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locationResponse.Next
	if locationResponse.Previous != nil {
		cfg.Previous = *locationResponse.Previous
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResponse, err := pokeapi.GetLocationAreas(cfg.Previous)
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
