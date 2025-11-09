package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

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

func getLocationAreaResponse(url string) (LocationAreaResponse, error) {
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

func printLocationArea(resp LocationAreaResponse) {
	data := resp.Results
	results := []string{}
	for _, location := range data {
		results = append(results, location.Name)
	}
	fmt.Print(strings.Join(results, "\n") + "\n")
}

func updatePagination(pagination *PaginationConfig, resp LocationAreaResponse) {
	pagination.Back = resp.Previous
	pagination.Next = resp.Next
}
