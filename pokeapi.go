package main

import (
	"net/http"
)

type Response struct {
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
	http.Get(url)
}
