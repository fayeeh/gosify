package main

import (
	"errors"
	"net"
	"os"
	"time"
)

func main() {
	port := ":3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	server := NewServer(port)
	cache := NewCache()
	cache.Run()

	server.AddCommands(
		&Command{
			Name:    "help",
			Aliases: []string{"h", "HELP"},
			Run: func(server *Server, args []string, conn net.Conn) error {
				helpmsg := `Usage: 
[]: optional, <>: required
commands:
  help shows this command  help
  echo echo                echo <msg>
  set set data			   set <key> <value> [duration] # Look https://pkg.go.dev/time#ParseDuration for duration
  get get data             get <key>
  rm  delete data          rm <key>
`

				Write(conn, helpmsg)
				return nil
			},
		},
		&Command{
			Name:    "set",
			Aliases: []string{"SET"},
			Run: func(server *Server, args []string, conn net.Conn) error {
				if len(args) < 3 {
					return errors.New("err: too few arguments for echo command\n")
				}

				if len(args) == 4 {
					dr, err := time.ParseDuration(args[3])
					if err != nil {
						return err
					}

					now := time.Now()
					cache.Add(args[1], args[2], now.Add(dr))
				} else {
					cache.Add(args[1], args[2], time.Time{})
				}

				Write(conn, "ok\n")
				return nil
			},
		},
		&Command{
			Name:    "get",
			Aliases: []string{"GET"},
			Run: func(server *Server, args []string, conn net.Conn) error {
				if len(args) < 2 {
					return errors.New("err: too few arguments for get command\n")
				}

				item := cache.Get(args[1])

				if item == nil {
					Write(conn, "null\n")
					return nil
				}

				Write(conn, item.Value+"\n")

				return nil
			},
		},
		&Command{
			Name:    "remove",
			Aliases: []string{"rm", "delete", "RM", "DELETE"},
			Run: func(server *Server, args []string, conn net.Conn) error {
				if len(args) < 2 {
					return errors.New("err: too few arguments for remove command\n")
				}

				cache.Remove(args[1])
				Write(conn, "ok\n")

				return nil
			},
		},
	)

	server.Start()
}

func getDuration(dr string) time.Duration {
	switch dr {
	case "-S":
		return time.Second
	case "-M":
		return time.Minute
	case "-H":
		return time.Hour
	default:
		return 1
	}
}
