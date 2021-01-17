package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	Helper "./Helper"
	Structure "./Structure"
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
	//Converted into the command data structure
	command := Structure.NewCmd(string(cmd), string(args), c.conn)

	//Routing the command to the right handler function
	switch string(cmd) {
	case "LOGIN":
		c.login(command)
	default:
		c.conn.Write([]byte("Send a recognizable command to the TCP Server"))
	}
}

func (c *client) login(command *Structure.Command) {
	//Processing body to get the right arguments
	args := Helper.ExtractingArgumentsFromCommands("LOGIN", command.Body)
	//Checking the validity of the credentials
	check := Helper.ValidateLogin(args["username"], args["password"])
	//If user is authenticated, get a bearer token and return it to the HTTP Server
	if check == true {
		token := Helper.CreateToken(args["username"])
		Helper.SendToHTTPServer(c.conn, "Ok, logged in with token: "+token)
	} else {
		token := "Invalid Credentials"
		Helper.SendToHTTPServer(c.conn, "Ok, logged in with token: "+token)
	}
}
