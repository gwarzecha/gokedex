package pokeapi

import (
	"net/http"
	"time"

	"github.com/gwarzecha/gokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

const baseURL = "https://pokeapi.co/api/v2/"

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
