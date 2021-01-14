package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

var (
	DELIMITER = []byte(`\r\n`)
)

type client struct {
	conn     net.Conn
	outbound chan<- command
	username string
}

func newClient(conn net.Conn, o chan<- command, username string) *client {
	return &client{
		conn:     conn,
		outbound: o,
		username: username,
	}
}

func (c *client) read() error {
	for {
		msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err != nil {
			return err
		}
		c.handle(msg)
	}
	return nil
}

func (c *client) handle(message []byte) {
	fmt.Println("Handling message...")
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

	switch string(cmd) {
	case "LOGIN":
		c.login(cmd, args)
	default:
		fmt.Println("In default logic...")
		c.login(cmd, args)
	}
}

func (c *client) login(cmd []byte, args []byte) {
	c.outbound <- command{
		conn:   c.conn,
		id:     cmd,
		sender: c.username,
		body:   args,
	}
}
