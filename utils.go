package main

import (
	"bufio"
	"net"
)

func Read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return str, nil
}

func Write(conn net.Conn, msg string) error {
	writer := bufio.NewWriter(conn)
	_, err := writer.WriteString(msg)

	if err == nil {
		writer.Flush()
	}

	return err
}
