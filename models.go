package main

import "github.com/zipsonic/pokedexcli/api"

const BaseAPIURL string = "https://pokeapi.co/api/v2/"
const BaseLocationURL string = BaseAPIURL + "location-area/"
const BasePokemonURL string = BaseAPIURL + "pokemon/"

type Config struct {
	Previous   string
	Next       string
	Areas      []string
	Encounters []string
	Caught     map[string]api.PokemonResponse
}

type Command struct {
	name string
	desc string
	cmd  func(*Config, []string) error
}
