package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const mapRequestLimit = 20

var commandRegistry = map[string]*cliCommand{}

func init() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
}

func startRepl() {
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		command := cleanedInput[0]
		commandInfo, ok := commandRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		commandInfo.Callback()
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
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
	// TODO: incorporate PokeAPI and mapRequestLimit
	http.Get("blah")
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
