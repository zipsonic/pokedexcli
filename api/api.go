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

func GetLocationArea(url string) LocationAreaResponse {

	body := getPokeData(url)

	var locationAreaResponse LocationAreaResponse
	if err := json.Unmarshal(body, &locationAreaResponse); err != nil {
		fmt.Println("Error Unmarshalling Data")
	}

	return locationAreaResponse
}

func GetExploreArea(url string) ExploreAreaResponse {

	body := getPokeData(url)

	var exploreAreaResponse ExploreAreaResponse
	if err := json.Unmarshal(body, &exploreAreaResponse); err != nil {
		fmt.Println("Error Unmarshalling Data")
	}

	return exploreAreaResponse
}

func GetPokemon(url string) PokemonResponse {

	body := getPokeData(url)

	var pokemonResponse PokemonResponse
	if err := json.Unmarshal(body, &pokemonResponse); err != nil {
		fmt.Println("Error Unmarshalling Data")
	}

	return pokemonResponse
}
