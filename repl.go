package main

import (
	"fmt"
	"os"
)

func startRepl() {
	// TODO: import functionality from main.go
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, commandInfo := range commandRegistry {
		fmt.Printf("%v: %v\n", commandInfo.Name, commandInfo.Description)
	}
	return nil
}

func commandMap() error {
	// TODO: refactor into repl.go
	return nil
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func registerCommand(name, description string, callback func() error) {
	commandRegistry[name] = &cliCommand{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}
