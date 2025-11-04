package main

import (
	"fmt"
)

func commandHelp() error {
	helpString := ""
	for _, value := range commandsRegistry {
		helpString += fmt.Sprintf("%s: %s\n", value.name, value.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", helpString)
	return nil
}
