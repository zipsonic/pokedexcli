package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/zipsonic/pokedexcli/api"
)

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

func wrappedCmdMap(config *Config, url string) error {

	if url == "" {
		url = BaseLocationURL + "?offset=0&limit=20"
	}

	returnVal := api.GetLocationArea(url)

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

	url := BaseLocationURL + area

	//Retrieve JSON for explored Area
	returnVal := api.GetExploreArea(url)
	//Reset Pokemon encounters - Only able to catch pokemon in currently explored area
	config.Encounters = nil

	fmt.Printf("Exploring %s\n", area)
	fmt.Println("Pokemon found:")

	//List Pokemon Encountered and add to config for possible catch
	for _, pokemon := range returnVal.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		config.Encounters = append(config.Encounters, pokemon.Pokemon.Name)
	}

	return nil
}

func cmdCatch(config *Config, cmdSlice []string) error {

	// if there are no arguments return error
	if len(cmdSlice) == 1 {
		fmt.Println("Catching.... a cold?")
		return nil
	}

	// Don't allow a pokemon to be caught if not found in the current area
	if !slices.Contains(config.Encounters, cmdSlice[1]) {
		fmt.Println("You don't see that Pokemon")
		return nil
	}

	url := BasePokemonURL + cmdSlice[1]

	//Retrieve JSON for specified pokemon
	returnVal := api.GetPokemon(url)

	//Based on random sample of base Experience, this should result in a 25% to 35% catch failure
	caughtIndex := rand.Intn(250)

	//Simulate Pokeball actions
	fmt.Printf("Throwing Pokeball at %s...\n", returnVal.Name)
	time.Sleep(1 * time.Second)
	fmt.Println("Wiggle....")
	time.Sleep(2 * time.Second)
	fmt.Println("Wiggle........")
	time.Sleep(3 * time.Second)

	if caughtIndex > returnVal.BaseExperience {
		fmt.Println("CLICK!")
		//fmt.Printf("CI: %d BE: %d\n", caughtIndex, returnVal.BaseExperience)
		fmt.Printf("You caught %s!!!\n", returnVal.Name)
		config.Caught[returnVal.Name] = returnVal
	} else {
		fmt.Printf("%s escaped!\n", returnVal.Name)
	}

	return nil
}

func cmdHelp(config *Config, cmdSlice []string) error {

	// if there are no arguments return help description
	if len(cmdSlice) == 1 {
		fmt.Println(cliCommand["help"].desc)
		return nil
	}

	if cmdSlice[1] == "list" {
		for _, cmd := range cliCommand {
			fmt.Printf("%s - %s\n\n", cmd.name, cmd.desc)
		}
		return nil
	}

	// Return command description, if command found
	if cmd, ok := cliCommand[cmdSlice[1]]; ok {
		fmt.Println(cmd.desc)
	} else {
		fmt.Println("Command not found")
	}

	return nil
}

func cmdExit(config *Config, cmdSlice []string) error {
	fmt.Println("Exiting Pokedex...")
	fmt.Println("Goodbye.")
	os.Exit(0)
	return nil
}
