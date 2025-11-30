package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ajasf444/pokedexcli/internal/pokeapi"
)

// TODO: make use of this constant
const mapRequestLimit = 20

type Config struct {
	client     pokeapi.Client
	pagination pokeapi.Pagination
}

var (
	commandRegistry = map[string]*cliCommand{}
	pagination      = &pokeapi.Pagination{
		Next: "https://pokeapi.co/api/v2/location-area",
		Back: "",
	}
)

func init() {
	registerCommand("exit", "Exits the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
	registerCommand("map", "Displays 20 locations", commandMap)
	registerCommand("mapb", "Displays 20 previous locations", commandMapB)
}

func startRepl(cfg *Config) {
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}
		command := cleanedInput[0]
		commandInfo, ok := commandRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		commandInfo.Callback(pagination) // TODO: make use of the &cfg.pagination
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}

func commandExit(pagination *pokeapi.Pagination) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(pagination *pokeapi.Pagination) error {
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, commandInfo := range commandRegistry {
		fmt.Printf("%v: %v\n", commandInfo.Name, commandInfo.Description)
	}
	return nil
}

func commandMap(pagination *pokeapi.Pagination) error {
	url := pagination.Next
	results, err := pokeapi.GetLocationAreaResponse(url)
	if err != nil {
		return err
	}
	pokeapi.UpdatePagination(pagination, results)
	pokeapi.PrintLocationArea(results)
	return nil
}

func commandMapB(pagination *pokeapi.Pagination) error {
	url := pagination.Back
	if url == "" {
		fmt.Println("You are on the first page!")
		return nil
	}
	results, err := pokeapi.GetLocationAreaResponse(url)
	if err != nil {
		return err
	}
	pokeapi.UpdatePagination(pagination, results)
	pokeapi.PrintLocationArea(results)
	return nil
}

func commandExplore(pagination *pokeapi.Pagination, location string) error {
	// TODO: add logic
	return nil
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Pagination) error
}

func registerCommand(name, description string, callback func(*pokeapi.Pagination) error) {
	commandRegistry[name] = &cliCommand{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}
