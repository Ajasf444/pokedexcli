package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Ajasf444/pokedexcli/internal/pokecache"
)

// TODO: refactor most of this code to be more concise

// TODO: use a pokeapiClient for the caching
var cache = pokecache.NewCache(5 * time.Second)

func getRequest(url string) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, errors.New("error: failed to get response")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("error: failed to read response body")
	}
	return body, nil
}

func convertDataToLocationAreaResponse(data []byte) (LocationAreaResponse, error) {
	response := LocationAreaResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return LocationAreaResponse{}, errors.New("error: unable to unmarshal LocationAreaResponse")
	}
	return response, nil
}

func GetLocationAreaResponse(url string) (LocationAreaResponse, error) {
	data, ok := cache.Get(url)
	if !ok {
		var err error
		data, err = getRequest(url)
		if err != nil {
			return LocationAreaResponse{}, err
		}
	}
	cache.Add(url, data)
	jsonData, err := convertDataToLocationAreaResponse(data)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return jsonData, nil
}

func (c *Client) GetLocations(url string) (LocationAreaResponse, error) {
	if data, ok := c.cache.Get(url); ok {
		response := LocationAreaResponse{}
		if err := json.Unmarshal(data, &response); err != nil {
			return response, err
		}
		return response, nil
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, errors.New("error: failed to get response")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, errors.New("error: failed to read response body")
	}
	response := LocationAreaResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return LocationAreaResponse{}, errors.New("error: unable to unmarshal LocationAreaResponse")
	}
	return response, nil
}

func PrintLocationArea(resp LocationAreaResponse) {
	data := resp.Results
	for _, location := range data {
		fmt.Println(location.Name)
	}
}

func UpdatePagination(pagination *Pagination, resp LocationAreaResponse) {
	pagination.Next = resp.Next
	pagination.Back = resp.Previous
}
