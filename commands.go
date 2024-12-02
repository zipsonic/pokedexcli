package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zipsonic/pokedexcli/api"
)

func isNumberInRange(s string, min, max int) bool {
	// Convert the string to an integer
	num, err := strconv.Atoi(s)
	if err != nil {
		// If parsing fails, the string is not a valid number
		return false
	}

	// Check if the number is within the specified range
	return num >= min && num <= max
}

func cmdMap(config *Config, cmdSlice []string) error {

	if config.Next == "End" {
		fmt.Println("End of Location Areas")
		return nil
	}

	return wrappedCmdMap(config, config.Next)

}

func cmdMapb(config *Config, cmdSlice []string) error {

	if config.Previous == "Start" {
		fmt.Println("Already at Start of Location Areas")
		return nil
	}

	return wrappedCmdMap(config, config.Previous)

}

func wrappedCmdMap(config *Config, direction string) error {

	if direction == "" {
		direction = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

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

	config.Areas = nil

	for num, result := range returnVal.Results {
		fmt.Printf("%2d %s\n", (num + 1), result.Name)
		config.Areas = append(config.Areas, result.Name)
	}

	return nil

}

func cmdExplore(config *Config, cmdSlice []string) error {

	// if there are no arguments return error
	if len(cmdSlice) == 1 {
		fmt.Println("Must Specify an Area")
		return nil
	}

	//declaring area to use outside for condition
	area := ""

	// check to see if 1st argument is a string representation of a number, then return the index from the last map command
	// otherwise, find the area requested
	if isNumberInRange(cmdSlice[1], 1, 20) {
		index, _ := strconv.Atoi(cmdSlice[1])
		index -= 1
		area = config.Areas[index]
	} else {
		area = cmdSlice[1]
	}

	url := "https://pokeapi.co/api/v2/location-area/" + area

	returnVal := api.GetExploreArea(url)

	fmt.Printf("Exploring %s\n", area)
	fmt.Println("Pokemon found:")

	for _, pokemon := range returnVal.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func cmdHelp(config *Config, cmdSlice []string) error {
	fmt.Println(cliCommand["help"].desc)
	return nil
}

func cmdExit(config *Config, cmdSlice []string) error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}
