package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zipsonic/pokedexcli/api"
	"github.com/zipsonic/pokedexcli/pokecache"
)

var cliCommand = make(map[string]Command)

func registerCommand(name, desc string, cmd func(*Config, []string) error) {
	cliCommand[name] = Command{name: name, desc: desc, cmd: cmd}
}

func init() {
	registerCommand("help", "\"help [command]\" will give you a description of the command\n\"help list\" will list all commands.", cmdHelp)
	registerCommand("exit", "Exits the Pokedex", cmdExit)
	registerCommand("map", "Returns a page of location Data", cmdMap)
	registerCommand("mapb", "Returns prior page of location Data", cmdMapb)
	registerCommand("explore", "Explore areas returned from map & mapb. Calls require either area name or number.", cmdExplore)
	registerCommand("catch", "Attempt to catch Pokemon in Explored area. Must specify pokemon to catch", cmdCatch)
	registerCommand("inspect", "Returns the stats of the specified caught Pokemon", cmdInspect)
}

func main() {

	pokecache.NewCache(20 * time.Minute)
	caught := make(map[string]api.PokemonResponse)
	config := Config{"", "", nil, nil, caught}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Pokedex > ")

	for scanner.Scan() {

		cmdSlice := strings.Split(scanner.Text(), " ")

		if cmdExec, ok := cliCommand[cmdSlice[0]]; ok {

			cmdExec.cmd(&config, cmdSlice)
		}
		fmt.Print("Pokedex > ")
	}
}
