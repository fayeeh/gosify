package main

type CallbackFunction = func([]string) bool

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

func NewServer(port string) *Server {
	return &Server{
		commands: make([]*Command, 0),
		Port:     port,
	}
}
