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

var cliCommand = make(map[string]Command)

func registerCommand(name, desc string, cmd func() error) {
	cliCommand[name] = Command{name: name, desc: desc, cmd: cmd}
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

func init() {
	registerCommand("help", "Prints this help screen", cmdHelp)
	registerCommand("exit", "Exits the Pokedex", cmdExit)
	registerCommand("map", "Returns a page of location Data", cmdMap)
	registerCommand("mapb", "Returns prior page of location Data", cmdMapb)
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
