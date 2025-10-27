package main

import (
	"net/http"
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

func getRequest(url string) {
	// TODO: use PokeAPI to make get request
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		// TODO: error
	}
	defer resp.Body.Close()
}
