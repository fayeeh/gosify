package main

import (
	"os"
	"fmt"
)

func main() {
	port := ":3000" 
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Println(port)
	server := NewServer(port)
	
	server.AddCommands(
		&Command{
			Name: "help",
			Description: "Shows this message",
			Run: func(server *Server, args []string) {
			},
		},
	)
	
	server.Start()
}
