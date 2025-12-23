package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) getPokemon(name string) (Pokemon, error) {
	url := baseURL + "/ability/" + name
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, errors.New("error: unable to generate request")
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, errors.New("error: failed to get response")
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, errors.New("error: failed to read response body")
		}
		c.cache.Add(url, data)
	}
	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, nil
	}
	return pokemon, nil
}

func (c *Client) CatchPokemon(name string) error {
	pokemon, err := c.getPokemon(name)
	if err != nil {
		return err
	}
	print(pokemon)
	//TODO: based on base XP generate whether Pokemon was caught
	//TODO: print capture or escape
	//TODO: add caught pokemon to pokedex
	return nil
}
