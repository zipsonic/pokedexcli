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

func getPokeData(url string) []byte {

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

	return body
}

func fetch[T any](url string) T {
	body := getPokeData(url)
	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error Unmarshalling Data")
	}
	return result
}

func GetLocationArea(url string) LocationAreaResponse {
	return fetch[LocationAreaResponse](url)
}

func GetExploreArea(url string) ExploreAreaResponse {
	return fetch[ExploreAreaResponse](url)
}

func GetPokemon(url string) PokemonResponse {
	return fetch[PokemonResponse](url)
}
