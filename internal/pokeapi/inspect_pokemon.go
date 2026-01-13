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
	// TODO: add logic
}
