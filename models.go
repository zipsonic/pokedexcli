package main

type Config struct {
	Previous string
	Next     string
	Areas    []string
}

type Command struct {
	name string
	desc string
	cmd  func(*Config, []string) error
}
