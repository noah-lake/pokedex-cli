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

var commandsRegistry map[string]cliCommand

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	helpString := ""
	for _, value := range commandsRegistry {
		helpString += fmt.Sprintf("%s: %s\n", value.name, value.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", helpString)
	return nil
}

func init() {
	commandsRegistry = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		inputString := input.Text()
		cleanedInput := cleanInput(inputString)
		command, ok := commandsRegistry[cleanedInput[0]]
		if !ok {
			fmt.Print("Unknown command")
		} else {
			command.callback()
		}
	}
}
