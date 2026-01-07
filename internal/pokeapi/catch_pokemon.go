package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const maxXP float64 = 608

func (c *Client) getPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name + "/"
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
		data, err = io.ReadAll(resp.Body) // walrus operator created a new "data" variable in block that was not passed out for unmarshalling
		if err != nil {
			return Pokemon{}, errors.New("error: failed to read response body")
		}
		c.cache.Add(url, data)
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, errors.New("error: unable to unmarshal Pokemon")
	}
	return pokemon, nil
}

func throwPokeball(name string) {
	fmt.Println("Throwing a Pokeball at " + name + "...")
}

func (c *Client) CatchPokemon(name string) error {
	pokemon, err := c.getPokemon(name)
	if err != nil {
		return err
	}
	pokemonName := cases.Title(language.English).String(pokemon.Name)
	throwPokeball(pokemonName)
	// TODO: based on base XP generate whether Pokemon was caught
	num := rand.Intn(pokemon.BaseExperience)
	caught := num > pokemon.BaseExperience/4
	if !caught {
		fmt.Println(pokemonName + " escaped!")
		return nil
	}

	c.caughtPokemon[name] = pokemon
	fmt.Println(pokemonName + " was caught!")
	return nil
}

func (c *Client) CatchPokemonSimple(name string) error {
	pokemon, err := c.getPokemon(name)
	if err != nil {
		return err
	}
	pokemonName := pokemon.Name
	throwPokeball(pokemonName)
	num := rand.Intn(pokemon.BaseExperience)
	caught := num > pokemon.BaseExperience/4
	if !caught {
		fmt.Println(pokemonName + " escaped!")
		return nil
	}

	c.caughtPokemon[name] = pokemon
	fmt.Println(pokemonName + " was caught!")
	return nil
}
