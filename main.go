package main

import (
	"time"

	"github.com/Turocz3k/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
