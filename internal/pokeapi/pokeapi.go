package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// TODO: refactor most of this code to be more concise

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
