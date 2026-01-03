package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{}

	cfg := &config{}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Show location areas (next page)",
		callback: func() error {
			return commandMap(cfg)
		},
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Show location areas (previous page)",
		callback: func() error {
			return commandMapBack(cfg)
		},
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp(commands),
	}

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		if cmd, ok := commands[cmdName]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Printf("Error executing %s: %v\n", cmdName, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
