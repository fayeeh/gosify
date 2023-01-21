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
	server := NewServer(port)

	fmt.Println(server.Port)
}
