package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/zipsonic/pokedexcli/pokecache"
)

var cliCommand = make(map[string]Command)

func registerCommand(name, desc string, cmd func(*Config) error) {
	cliCommand[name] = Command{name: name, desc: desc, cmd: cmd}
}

func init() {
	registerCommand("help", "Prints this help screen", cmdHelp)
	registerCommand("exit", "Exits the Pokedex", cmdExit)
	registerCommand("map", "Returns a page of location Data", cmdMap)
	registerCommand("mapb", "Returns prior page of location Data", cmdMapb)
}

func main() {

	pokecache.NewCache(10 * time.Second)
	config := Config{"", ""}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex > ")

	for scanner.Scan() {

		cmdInput := scanner.Text()

		if cmdExec, ok := cliCommand[cmdInput]; ok {

			cmdExec.cmd(&config)
		}
		fmt.Print("Pokedex > ")
	}
}
