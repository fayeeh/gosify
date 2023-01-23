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
			Run: func(server *Server, args []string, conn net.Conn) error {
				helpmsg := `Usage: 
[]: optional, <>: required
commands:
  help shows this command  help
  echo echo                echo <msg>`

				Write(conn, helpmsg)
				return nil
			},
		},
		&Command{
			Name:        "echo",
			Aliases:     []string{},
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
