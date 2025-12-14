package main

import (
	"time"

	"github.com/Ajasf444/pokedexcli/internal/pokeapi"
)

// TODO: move baseURL to another file
func main() {
	cfg := &Config{
		client:     pokeapi.NewClient(5*time.Minute, 5*time.Second),
		pagination: &pokeapi.Pagination{},
	}
	startRepl(cfg)
}
