package main

import (
	"bufio"
	"bytes"
	"encoding/json"
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
	case "SHOW_PROFILE":
		c.showProfile(command)
	case "LOGOUT":
		c.logout(command)
	default:
		Helper.SendToHTTPServer(c.conn, "Send a recognizable command to the TCP Server")
	}
}

func (c *client) logout(command *Structure.Command) {
	args := Helper.ExtractingArgumentsFromCommands("LOGIN", command.Body)
	tokenAuth := args["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	deleted, delErr := Helper.DeleteAuth(tokenAuthMap["AccessUUID"])
	if delErr != nil || deleted == 0 { //if any goes wrong
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
	}
	Helper.SendToHTTPServer(c.conn, "Logged out!")
}

func (c *client) showProfile(command *Structure.Command) {
	args := Helper.ExtractingArgumentsFromCommands("LOGIN", command.Body)
	tokenAuth := args["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)

	username, err := Helper.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
	}
	//Function to display the profile of the user from the database
	Helper.SendToHTTPServer(c.conn, "Welcome "+username)
}

func (c *client) login(command *Structure.Command) {
	//Processing body to get the right arguments
	args := Helper.ExtractingArgumentsFromCommands("LOGIN", command.Body)
	//Checking the validity of the credentials
	check := Helper.ValidateLogin(args["username"], args["password"])
	//If user is authenticated, get a bearer token and return it to the HTTP Server
	if check == true {
		//Token is created for auth
		token, _ := Helper.CreateToken(args["username"])
		//User is saved in redis
		saveErr := Helper.CreateAuth(args["username"], token)
		if saveErr != nil {
			token := "Invalid Credentials"
			Helper.SendToHTTPServer(c.conn, token)
		}
		//Data prepared for sending to HTTP server
		tokens := map[string]string{
			"access_token": token.AccessToken,
		}
		out, _ := json.Marshal(tokens)
		//Data sent to HTTP server
		Helper.SendToHTTPServer(c.conn, string(out))
	} else {
		token := "Invalid Credentials"
		Helper.SendToHTTPServer(c.conn, token)
	}
}
