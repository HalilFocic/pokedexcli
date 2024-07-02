package main

import "fmt"

func callbackHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for _, v := range commands {
		m := fmt.Sprintf("%s: %s \n", v.name, v.description)
		fmt.Print(m)
	}
	fmt.Print("\n")
	return nil
}
