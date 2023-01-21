package main

import (
	"net"
)

type CallbackFunction = func(*Server, []string)

type Command struct {
	Name        string
	Description string
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
func (s *Server) Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// TODO handle error
	}

	for { 
		conn, err := ln.Accept()
		if err != nil {
			//TODO Handle Error
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
}
func NewServer(port string) *Server {
	return &Server{
		commands: make([]*Command, 0),
		Port:     port,
	}
}
