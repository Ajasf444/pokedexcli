package main

import (
	"time"

	"github.com/Ajasf444/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &Config{
		client:     pokeapi.NewClient(5*time.Minute, 5*time.Second),
		pagination: &pokeapi.Pagination{Next: "https://pokeapi.co/api/v2/location-area"},
	}
	startRepl(cfg)
}
