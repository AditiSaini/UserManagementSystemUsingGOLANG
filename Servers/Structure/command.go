package structure

import "net"

func NewCmd(id string, body map[string]string, conn net.Conn) *Command {
	return &Command{
		Id:   id,
		Body: body,
		Conn: conn,
	}
}

type Command struct {
	//Identification of the commands
	Id string
	//Body of the command sent by the sender
	Body map[string]string
	//Connection being used to connect to the server
	Conn net.Conn
}
