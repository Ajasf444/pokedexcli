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

var cache = pokecache.NewCache(5 * time.Second)

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaContent struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func PrintLocationArea(resp LocationAreaResponse) {
	data := resp.Results
	for _, location := range data {
		fmt.Println(location.Name)
	}
}

func UpdatePagination(pagination *PaginationConfig, resp LocationAreaResponse) {
	pagination.Next = resp.Next
	pagination.Back = resp.Previous
}
