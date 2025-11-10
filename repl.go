package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const mapRequestLimit = 20

var (
	commandRegistry = map[string]*cliCommand{}
	pagination      = &PaginationConfig{
		Next: "",
		Back: "",
	}
)

func init() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
	registerCommand("map", "Displays 20 locations", commandMap)
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
		commandInfo.Callback(pagination)
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}

func commandExit(pagination *PaginationConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(pagination *PaginationConfig) error {
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, commandInfo := range commandRegistry {
		fmt.Printf("%v: %v\n", commandInfo.Name, commandInfo.Description)
	}
	return nil
}

func commandMap(pagination *PaginationConfig) error {
	url := "https://pokeapi.co/api/v2/location-area/" // TODO: check pagination
	results, err := getLocationAreaResponse(url)
	if err != nil {
		return err
	}
	updatePagination(pagination, results)
	printLocationArea(results)
	return nil
}

func commandMapb(pagination *PaginationConfig) error {
	// TODO: incorporate pokeapi.go getRequest()
	return nil
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*PaginationConfig) error
}

func registerCommand(name, description string, callback func(*PaginationConfig) error) {
	commandRegistry[name] = &cliCommand{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}
