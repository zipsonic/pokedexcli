package main

type Config struct {
	Previous string
	Next     string
}

type Command struct {
	name string
	desc string
	cmd  func(*Config) error
}
