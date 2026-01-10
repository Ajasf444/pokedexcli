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
	fmt.Println(pokemon) // TODO: handle pokemon data
	return nil
}
