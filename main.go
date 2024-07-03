package main

import "github.com/HalilFocic/pokedexcli/internal/pokeapi"

type config struct {
	pokeApiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
    cfg := config {
        pokeApiClient: pokeapi.NewClient(),
    }
	startRepl(&cfg)
}
