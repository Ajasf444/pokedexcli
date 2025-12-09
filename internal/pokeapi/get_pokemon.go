package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// TODO: update this signature
func (c *Client) GetLocationContent(pageURL *string) (LocationAreaContent, error) {
	url := baseURL + "/location-area" // TODO: update this url
	if pageURL != nil {
		url = *pageURL
	}
	data, ok := c.cache.Get(url)
	if !ok {
		resp, err := c.httpClient.Get(url)
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

func PrintPokemon(resp LocationAreaResponse) {
	data := resp.Results
	for _, location := range data {
		fmt.Println(location.Name)
	}
}
