package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
    if len(cfg.caughtPokemon) == 0 {
        return fmt.Errorf("You have not caught any pokemons!")
    }
    for pokemonName := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n",pokemonName)
    }
    return nil

}
