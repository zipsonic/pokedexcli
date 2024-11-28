package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationArea(url string) LocationAreaResponse {

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(resp.Status)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Reading Body")
	}

	var locationAreaResponse LocationAreaResponse
	if err := json.Unmarshal(body, &locationAreaResponse); err != nil {
		fmt.Println("Error Unmarshalling Data")
	}

	return locationAreaResponse
}
