package helper

import (
	"bufio"
	"fmt"
	"net"
	Connection "servers/ConnectionPool"
)

func SendToHTTPServer(conn net.Conn, response string) {
	//Connecting to the HTTP Server to send the response
	conn.Write([]byte(response))
	conn.Close()
}

func GetResponseFromTCPServer(command string, c net.Conn, pool *Connection.GncpPool) string {
	//Text is the command to be sent to the TCP server
	text := command
	fmt.Fprintf(c, text+"\n")

	//Receiving message from the TCP server
	message, _ := bufio.NewReader(c).ReadString('\n')
	Connection.CloseTCPConnection(c, pool)
	return message
}
