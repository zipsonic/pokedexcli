package main

import (
	"fmt"
	"os"
)

func cmdMap() error {
	return nil
}

func cmdMapb() error {
	return nil
}

func cmdHelp() error {
	fmt.Println(cliCommand["help"].desc)
	//Is using the struct desc the best method here?
	return nil
}

func cmdExit() error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}
