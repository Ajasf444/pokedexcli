package pokeapi

import "fmt"

func (c *Client) InspectPokemon(name string) error {
	pokemon, ok := c.caughtPokemon[name]
	if !ok {
		fmt.Printf("%v has not been caught!", name)
	}
	fmt.Println(pokemon) // TODO: handle pokemon data
	return nil
}
