package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zipsonic/pokedexcli/pokecache"
)

var cliCommand = make(map[string]Command)

func registerCommand(name, desc string, cmd func(*Config, []string) error) {
	cliCommand[name] = Command{name: name, desc: desc, cmd: cmd}
}

func init() {
	registerCommand("help", "Prints this help screen", cmdHelp)
	registerCommand("exit", "Exits the Pokedex", cmdExit)
	registerCommand("map", "Returns a page of location Data", cmdMap)
	registerCommand("mapb", "Returns prior page of location Data", cmdMapb)
	registerCommand("explore", "Explore areas returned from map & mapb. Calls require either area name or number.", cmdExplore)
}

func main() {

	pokecache.NewCache(10 * time.Second)
	config := Config{"", "", nil}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex > ")

	for scanner.Scan() {

		cmdInput := scanner.Text()

		cmdSlice := strings.Split(cmdInput, " ")

		if cmdExec, ok := cliCommand[cmdSlice[0]]; ok {

			cmdExec.cmd(&config, cmdSlice)
		}
		fmt.Print("Pokedex > ")
	}
}
