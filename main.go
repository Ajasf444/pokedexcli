package main

import (
	"strings"
)

func main() {
	startRepl()
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
