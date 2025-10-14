package main

import (
	"strings"
)

var commandRegistry = map[string]*cliCommand{}

func init() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
}

func main() {
	startRepl()
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
