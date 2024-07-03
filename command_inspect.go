package main

import (
	"errors"
	"fmt"
)

func callBackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon provided")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught %s yet", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
    fmt.Print("Types: ")
	for _, typ := range pokemon.Types {
		fmt.Printf(" %s", typ.Type.Name)
	}
    fmt.Print("\n")
	return nil

}
