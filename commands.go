package main

import (
	"fmt"
	"os"

	"github.com/zipsonic/pokedexcli/api"
)

func cmdMap(config *Config) error {

	if config.Next == "End" {
		fmt.Println("End of Location Areas")
		return nil
	}

	returnVal := api.GetLocationArea(config.Next)

	if returnVal.NextPage == nil {
		config.Previous = *returnVal.PreviousPage
		config.Next = "End"
	} else {
		config.Previous = *returnVal.PreviousPage
		config.Next = *returnVal.NextPage
	}

	for _, result := range returnVal.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func cmdMapb(config *Config) error {

	if config.Next == "Begin" {
		fmt.Println("At Beginning of Location Areas")
		return nil
	}

	returnVal := api.GetLocationArea(config.Previous)

	if returnVal.PreviousPage == nil {
		config.Previous = "Begin"
		config.Next = *returnVal.NextPage
	} else {
		config.Previous = *returnVal.PreviousPage
		config.Next = *returnVal.NextPage
	}

	for _, result := range returnVal.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func cmdHelp(config *Config) error {
	fmt.Println(cliCommand["help"].desc)
	return nil
}

func cmdExit(config *Config) error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}
