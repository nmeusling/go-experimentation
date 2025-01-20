package main

import (
	"time"

	"github.com/nmeusling/go-experimentation/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokemon{},
	}
	startRepl(cfg)
}
