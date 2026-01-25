package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationContent(location string) (LocationAreaContent, error) {
	url := baseURL + "/location-area/" + location
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaContent{}, errors.New("error: failed to generate request")
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaContent{}, errors.New("error: failed to get response")
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaContent{}, errors.New("error: failed to read response body")
		}
		c.cache.Add(url, data)
	}
	response := LocationAreaContent{}
	if err := json.Unmarshal(data, &response); err != nil {
		return LocationAreaContent{}, errors.New("error: unable to unmarshal LocationAreaContent")
	}
	return response, nil
}

func PrintPokemon(resp LocationAreaContent) {
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}
