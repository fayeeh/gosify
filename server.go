package main

import (
	"net"
	"fmt"
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
	ln, err := net.Listen("tcp", s.Port)
	if err != nil {
		// TODO handle error
	}
	defer ln.Close()
	
	fmt.Println("Server started at: ", s.Port)

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
	
	fmt.Println("Client Connected")
	for {
		str, err := Read(conn)
		
		if err != nil {
			break
		}

		err = Write(conn, str) 

		if err != nil {
			break
		}
	}
	fmt.Println("Client disconnected")

}
func NewServer(port string) *Server {
	return &Server{
		commands: make([]*Command, 0),
		Port:     port,
	}
}
