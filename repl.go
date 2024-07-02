package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()

        fmt.Print(text)
	}
	

}

func cleanInput(str string) []string{
    lowered := strings.ToLower(str)
    words := strings.Fields(lowered)
    return words

}
func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
	return commands
}

func commandHelp() error {
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
func commandExit() error {
	os.Exit(0)
	return nil
}
