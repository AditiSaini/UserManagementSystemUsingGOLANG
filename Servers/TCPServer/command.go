package main

import "net"

type ID int

const (
	REG ID = iota
	LOGIN
	LOGOUT
	UPLOAD
	UPDATE
	GET
)

type command struct {
	//Identification of the commands
	id []byte
	//Sender of the commands identified by a username
	sender string
	//Body of the command sent by the sender
	body []byte
	conn net.Conn
}
