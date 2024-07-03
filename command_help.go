package main

import "fmt"

func callbackHelp(cfg *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for _, v := range commands {
		fmt.Printf("%s: %s \n", v.name, v.description)
	}
	fmt.Print("\n")
	return nil
}
