package pokeapi

import "fmt"

func (c *Client) Pokedex() error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf("  - %v\n", pokemon.Name)
	}
	return nil
}
