package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zipsonic/pokedexcli/pokecache"
)

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(time.Second * 10)
}

func GetLocationArea(url string) LocationAreaResponse {

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	var body []byte
	var ok bool

	body, ok = cache.Get(url)

	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(resp.Status)
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error Reading Body")
		}
		cache.Add(url, body)
	}

	var locationAreaResponse LocationAreaResponse
	if err := json.Unmarshal(body, &locationAreaResponse); err != nil {
		fmt.Println("Error Unmarshalling Data")
	}

	return locationAreaResponse
}
