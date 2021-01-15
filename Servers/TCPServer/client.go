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
	conn net.Conn
}

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
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
}

func (c *client) handle(message []byte) {
	fmt.Println("Handling command: " + string(message))

	//Processing commands sent by the HTTP Server
	//Step 1- Get the command
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	//Step 2- Extract the arguments of the command
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))

	//Routing the command to the right handler function
	switch string(cmd) {
	case "LOGIN":
		c.login(cmd, args)
	default:
		c.conn.Write([]byte("Send a recognizable command to the TCP Server"))
	}
}

func (c *client) login(cmd []byte, args []byte) {
	c.conn.Write([]byte("Ok, logged in!"))
	c.conn.Close()
}
