package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, parameters []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, parameters []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Print("\n\n")
	commands := getCommands()

	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}

func commandMap(cfg *config, parameters []string) error {
	locationResponse, err := cfg.pokeapiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locationResponse.Next
	if locationResponse.Previous != nil {
		cfg.Previous = *locationResponse.Previous
	}

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, parameters []string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationResponse, err := cfg.pokeapiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = cfg.Previous
	if locationResponse.Previous != nil {
		cfg.Previous = *locationResponse.Previous
	} else {
		cfg.Previous = ""
	}

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, parameters []string) error {
	location := ""
	if len(parameters) >= 1 {
		location = parameters[0]
	} else {
		fmt.Println("Please enter a location to explore")
		return nil
	}

	pokemonResponse, err := cfg.pokeapiClient.GetLocationPokemon(location)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + location + "...")
	fmt.Println("Pokemon found:")
	for _, pokemon := range pokemonResponse.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, parameters []string) error {
	pokemonName := ""
	if len(parameters) >= 1 {
		pokemonName = parameters[0]
	} else {
		fmt.Println("Please enter a pokemon to catch")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonData, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Println("Invalid pokemon, try again")
		return nil
	}
	baseExperience := pokemonData.BaseExperience
	catchTreshold := 50
	catchResult := rand.Intn(baseExperience) < catchTreshold
	if catchResult {
		fmt.Printf("You caught %s", pokemonName)
		pokeStats := pokemonStats{}
		pokeTypes := []string{}
		cfg.pokedex[pokemonName] = pokemon{
			pokemonData.Name,
			pokemonData.BaseExperience,
			pokemonData.Weight,
			pokemonData.Height,
			pokeStats,
			pokeTypes,
		}
	} else {
		fmt.Printf("%s escaped! Try again\n", pokemonName)
	}

	return nil
}

func commandInspect(cfg *config, parameters []string) error {
	pokemon := ""
	if len(parameters) >= 1 {
		pokemon = parameters[0]
	} else {
		fmt.Println("Please enter a pokemon to inspect")
		return nil
	}
	pokeData, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Inspecting %s\n", pokeData.name)
	fmt.Println(pokeData.height, pokeData.weight)
	return nil
}
