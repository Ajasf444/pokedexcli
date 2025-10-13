package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandRegistry = map[string]*cliCommand{}

func init() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
}

func main() {
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
