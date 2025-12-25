package pokeapi

import (
	"net/http"
	"time"

	"github.com/Ajasf444/pokedexcli/internal/pokecache"
)

type Client struct {
	cache         pokecache.Cache
	httpClient    http.Client
	caughtPokemon map[string]Pokemon
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		caughtPokemon: map[string]Pokemon{},
	}
}
