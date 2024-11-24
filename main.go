package main

import (
	"bufio"
	"fmt"
	"os"
)

type commands struct {
	name string
	desc string
	cmd  func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")

	for scanner.Scan() {
		fmt.Println("Pokedex >")
		cmdInput := scanner.Text()

	}
}
