package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) getPokemon(name string) (SimplePokemon, error) {
	url := baseURL + "/ability/" + name
	var data []byte
	var ok bool
	data, ok = c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return SimplePokemon{}, errors.New("error: unable to generate request")
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return SimplePokemon{}, errors.New("error: failed to get response")
		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return SimplePokemon{}, errors.New("error: failed to read response body")
		}
		c.cache.Add(url, data)
	}
	//TODO: try making a SimplePokemon struct with only BaseExperience
	pokemon := SimplePokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return SimplePokemon{}, errors.New("error: unable to unmarshal Pokemon")
	}
	return pokemon, nil
}

func (c *Client) CatchPokemon(name string) error {
	pokemon, err := c.getPokemon(name)
	if err != nil {
		return err
	}
	//TODO: based on base XP generate whether Pokemon was caught
	num := rand.Intn(pokemon.BaseExperience)
	caught := num > pokemon.BaseExperience/4
	if !caught {
		fmt.Println("Pokemon escaped!")
		return nil
	}

	c.caughtPokemon[name] = pokemon
	fmt.Println("Pokemon caught!")
	return nil
}
