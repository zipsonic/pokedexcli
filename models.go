package main

type Command struct {
	name string
	desc string
	cmd  func() error
}
