package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Ajasf444/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}

func commandExit(cfg *Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, arg string) error {
	fmt.Println("Usage:")
	fmt.Print("--------------------------------------\n")
	for _, commandInfo := range commandRegistry {
		fmt.Printf("%v: %v\n", commandInfo.Name, commandInfo.Description)
	}
	return nil
}

func commandMap(cfg *Config, arg string) error {
	url := cfg.pagination.Next
	results, err := cfg.client.GetLocations(url)
	if err != nil {
		return err
	}
	pokeapi.UpdatePagination(cfg.pagination, results)
	pokeapi.PrintLocationArea(results)
	return nil
}

func commandMapB(cfg *Config, arg string) error {
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
	if location == "" {
		return errors.New("Provide a location!")
	}
	results, err := cfg.client.GetLocationContent(location)
	if err != nil {
		return err
	}
	pokeapi.PrintPokemon(results)
	return nil
}

func commandCatch(cfg *Config, pokemon string) error {
	err := cfg.client.CatchPokemonSimple(pokemon)
	if err != nil {
		return err
	}
	return nil
}

func commandInspect(cfg *Config, pokemon string) error {
	return nil
}

func registerCommand(name, description string, callback func(*Config, string) error) {
	commandRegistry[name] = &cliCommand{
		Name:        name,
		Description: description,
		Callback:    callback,
	}
}
