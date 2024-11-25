package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	name string
	desc string
	cmd  func() error
}

var cliCommand = map[string]Command{
	"help": {
		name: "help",
		desc: "Prints this help screen",
		cmd:  cmdHelp,
	},
	"exit": {
		name: "exit",
		desc: "Exits the Pokedex",
		cmd:  cmdExit,
	},
	"map": {
		name: "map",
		desc: "Returns a page of location Data from the PokeAPI in 20 area increments",
		cmd:  cmdMap,
	},
	"mapb": {
		name: "mapb",
		desc: "Returns prior page of location Data from the PokeAPI in 20 area increments",
		cmd:  cmdMapb,
	},
}

func cmdMap() error {
	return nil
}

func cmdMapb() error {
	return nil
}

func cmdHelp() error {
	fmt.Println(cliCommand["help"].desc)
	return nil
}

func cmdExit() error {
	os.Exit(0)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex >")

	for scanner.Scan() {

		cmdInput := scanner.Text()

		if cmdExec, ok := cliCommand[cmdInput]; ok {

			cmdExec.cmd()
		}
		fmt.Print("Pokedex >")
	}
}
