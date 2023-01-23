package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type CallbackFunction = func(*Server, []string, net.Conn) error

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Run         CallbackFunction
}

type Server struct {
	commands []*Command
	Port     string
}

func (s *Server) AddCommands(cmds ...*Command) {
	s.commands = append(s.commands, cmds...)
}

func (s *Server) GetCommands() []*Command {
	return s.commands
}

func (s *Server) getCommand(name string) *Command {
	for _, cmd := range s.commands {
		if cmd.Name == name || contains(cmd.Aliases, name) {
			return cmd
		}
	}

	return nil
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.Port)
	if err != nil {
		// TODO handle error
	}
	defer ln.Close()

	fmt.Printf("Server started at: %s\n", s.Port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			//TODO Handle Error
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		str, err := Read(conn)

		if err != nil {
			break
		}

		args := strings.Fields(str)

		if len(args) < 1 {
			continue
		}

		if err = s.handleCommand(args, conn); err != nil {
			Write(conn, err.Error())
		}

	}

}

func (s *Server) handleCommand(args []string, conn net.Conn) error {
	cmd_name := args[0]

	if cmd := s.getCommand(cmd_name); cmd != nil {
		return cmd.Run(s, args, conn)
	}

	return errors.New("err: command not found\n")
}

func NewServer(port string) *Server {
	return &Server{
		commands: make([]*Command, 0),
		Port:     port,
	}
}

func contains[T comparable](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
