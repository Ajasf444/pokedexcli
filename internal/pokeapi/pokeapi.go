package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Ajasf444/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(10 * time.Second)

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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
	// TODO: add caching
	data, err := getRequest(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	jsonData, err := convertDataToLocationAreaResponse(data)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return jsonData, nil
}

func PrintLocationArea(resp LocationAreaResponse) {
	data := resp.Results
	results := []string{}
	for _, location := range data {
		results = append(results, location.Name)
	}
	fmt.Print(strings.Join(results, "\n") + "\n")
}

func UpdatePagination(pagination *PaginationConfig, resp LocationAreaResponse) {
	pagination.Back = resp.Previous
	pagination.Next = resp.Next
}
