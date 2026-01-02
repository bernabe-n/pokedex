package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)

	text = strings.ToLower(text)

	if text == "" {
		return []string{}
	}

	words := strings.Fields(text)

	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Help command: dynamically prints all commands
func commandHelp(commands map[string]cliCommand) func() error {
	return func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:\n")
		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}
}
