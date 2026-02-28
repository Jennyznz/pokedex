package pokeapi

import (
	"net/http"
	"time"
	"github.com/Jennyznz/pokedex/internal/pokecache"
)

type Client struct {
	HttpClient http.Client
	cache *pokecache.Cache
}

func NewClient(c *pokecache.Cache) Client {
	return Client {
		HttpClient: http.Client {
			Timeout: 10 * time.Second,
		},
		cache: c,
	}
}