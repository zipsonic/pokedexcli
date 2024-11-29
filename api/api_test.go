package api

import (
	"encoding/json"
	"testing"
)

const mockLocationAreaJSON = `
{
  "count": 1054,
  "next": "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
  "previous": null,
  "results": [
    { "name": "canalave-city-area", "url": "https://pokeapi.co/api/v2/location-area/1/" },
    { "name": "eterna-city-area", "url": "https://pokeapi.co/api/v2/location-area/2/" }
  ]
}`

func TestUnmarshalLocationAreaResponse(t *testing.T) {
	var locationAreaResponse LocationAreaResponse

	// Unmarshal the mock JSON
	err := json.Unmarshal([]byte(mockLocationAreaJSON), &locationAreaResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Validate fields
	if locationAreaResponse.Count != 1054 {
		t.Errorf("Expected Count to be 1054, got %d", locationAreaResponse.Count)
	}

	if *locationAreaResponse.Next != "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20" {
		t.Errorf("Expected Next to be 'https://pokeapi.co/api/v2/location-area/?offset=20&limit=20', got %s", *locationAreaResponse.Next)
	}

	if len(locationAreaResponse.Results) != 2 {
		t.Fatalf("Expected 2 results, got %d", len(locationAreaResponse.Results))
	}

	// Validate the first result
	firstResult := locationAreaResponse.Results[0]
	if firstResult.Name != "canalave-city-area" {
		t.Errorf("Expected first result name to be 'canalave-city-area', got %s", firstResult.Name)
	}

	if firstResult.Url != "https://pokeapi.co/api/v2/location-area/1/" {
		t.Errorf("Expected first result URL to be 'https://pokeapi.co/api/v2/location-area/1/', got %s", firstResult.Url)
	}
}
