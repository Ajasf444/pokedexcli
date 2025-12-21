package pokeapi

import (
	"encoding/json"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	url := baseURL + "/ability/" + name
	data, ok := c.cache.Get(url)
	if !ok {
		//TODO: client work
	}
	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, nil
	}
	return pokemon, nil
}
