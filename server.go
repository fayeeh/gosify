package main

type CallbackFunction = func([]string) bool

type Command struct {
	Name string
	Description string
	Run CallbackFunction
}

type Server struct {
	commands []*Command	
}
