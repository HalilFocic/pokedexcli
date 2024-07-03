package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("Failed to catch pokemon %s \n", pokemonName)
	}
	fmt.Printf("%s was caught! \n", pokemonName)
	return nil
}
