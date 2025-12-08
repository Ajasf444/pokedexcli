package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Ajasf444/pokedexcli/internal/pokeapi"
)

// TODO: make use of this constant
const mapRequestLimit = 20

type Config struct {
	client     pokeapi.Client
	pagination *pokeapi.Pagination
}

var commandRegistry = map[string]*cliCommand{}

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
		err := commandInfo.Callback(cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Usage:")
	fmt.Print("--------------------------------------\n")
	for _, commandInfo := range commandRegistry {
		fmt.Printf("%v: %v\n", commandInfo.Name, commandInfo.Description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	url := cfg.pagination.Next
	results, err := cfg.client.GetLocations(url)
	if err != nil {
		return err
	}
	pokeapi.UpdatePagination(cfg.pagination, results)
	pokeapi.PrintLocationArea(results)
	return nil
}

func commandMapB(cfg *Config) error {
	url := cfg.pagination.Back
	if url == nil {
		return errors.New("You are on the first page!")
	}
	results, err := cfg.client.GetLocations(url)
	if err != nil {
		return err
	}
	pokeapi.UpdatePagination(cfg.pagination, results)
	pokeapi.PrintLocationArea(results)
	return nil
}

func commandExplore(cfg *Config, location string) error {
	// TODO: add logic
	return nil
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func registerCommand(name, description string, callback func(*Config) error) {
	commandRegistry[name] = &cliCommand{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}
