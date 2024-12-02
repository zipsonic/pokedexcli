package main

import "github.com/zipsonic/pokedexcli/api"

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
