package main

import (
	"errors"
	"net"
	"os"
)

func main() {
	port := ":3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	server := NewServer(port)

	server.AddCommands(
		&Command{
			Name:        "help",
			Aliases:     []string{"h", "HELP"},
			Description: "Shows this message",
			Run: func(server *Server, args []string, conn net.Conn) error {
				Write(conn, "Hello\n")
				return nil
			},
		},
		&Command{
			Name:        "echo",
			Description: "Echo",
			Run: func(server *Server, args []string, conn net.Conn) error {
				if len(args) < 2 {
					return errors.New("err: too few arguments for echo command")
				}
				Write(conn, args[1]+"\n")
				return nil
			},
		},
	)

	server.Start()
}
