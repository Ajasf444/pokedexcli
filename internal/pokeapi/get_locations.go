package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to generate request")
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to get response")
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error: failed to read response body")
		}
		c.cache.Add(url, data)
	}
	response := LocationAreaResponse{}
	if err := json.Unmarshal(data, &response); err != nil {
		return LocationAreaResponse{}, errors.New("error: unable to unmarshal LocationAreaResponse")
	}
	return response, nil
}

func PrintLocationArea(resp LocationAreaResponse) {
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
}

func UpdatePagination(pagination *Pagination, resp LocationAreaResponse) {
	pagination.Next = resp.Next
	pagination.Back = resp.Previous
}
