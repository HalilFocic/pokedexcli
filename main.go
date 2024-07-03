package main

import (
	"time"
	"github.com/HalilFocic/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
    cfg := config {
        pokeApiClient: pokeapi.NewClient(time.Hour),
    }
	startRepl(&cfg)
}
