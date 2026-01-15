package pokeapi

import (
	"errors"
	"fmt"
)

func (c *Client) InspectPokemon(name string) error {
	pokemon, ok := c.caughtPokemon[name]
	if !ok {
		return errors.New(fmt.Sprintf("%v has not been caught!", name))
	}
	printPokemonStats(pokemon)
	return nil
}

func printPokemonStats(pokemon Pokemon) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %v\n", pokemonType.Type.Name)
	}
}
