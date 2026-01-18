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
	pagination *pokeapi.Pagination
}

var commandRegistry = map[string]*cliCommand{}

func init() {
	registerCommand("catch", "Attempt to catch the provided Pokemon", commandCatch)
	registerCommand("exit", "Exits the Pokedex", commandExit)
	registerCommand("explore", "Identify Pokemon in provided Location Area", commandExplore)
	registerCommand("help", "Displays a help message", commandHelp)
	registerCommand("inspect", "Inspects the provided Pokemon if caught", commandInspect)
	registerCommand("map", "Displays 20 locations", commandMap)
	registerCommand("mapb", "Displays 20 previous locations", commandMapB)
	registerCommand("pokedex", "Displays all caught Pokemon", commandPokedex)
}

func startRepl(cfg *Config) {
	var arg string
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
		if len(cleanedInput) >= 2 {
			arg = cleanedInput[1]
		}
		commandInfo, ok := commandRegistry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := commandInfo.Callback(cfg, arg)
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
