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

	return wrappedCmdMap(config, config.Next)

}

func cmdMapb(config *Config) error {

	if config.Previous == "Start" {
		fmt.Println("Already at Start of Location Areas")
		return nil
	}

	return wrappedCmdMap(config, config.Previous)

}

func wrappedCmdMap(config *Config, direction string) error {

	returnVal := api.GetLocationArea(direction)

	if returnVal.Previous == nil {
		config.Previous = "Start"
	} else {
		config.Previous = *returnVal.Previous
	}

	if returnVal.Next == nil {
		config.Next = "End"
	} else {
		config.Next = *returnVal.Next
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
