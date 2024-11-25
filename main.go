package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	name string
	desc string
	cmd  func()
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
}

func cmdHelp() {
	fmt.Println("Command Help Found")
}

func cmdExit() {
	fmt.Println("Command Exit found")
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
