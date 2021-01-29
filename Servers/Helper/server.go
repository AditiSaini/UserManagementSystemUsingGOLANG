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
}

func GetResponseFromTCPServer(command string, c net.Conn, pool *Connection.GncpPool) string {
	//Text is the command to be sent to the TCP server
	text := command
	fmt.Fprintf(c, text+"\n")

	//Receiving message from the TCP server
	message, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println("Error from the tcp server")
		fmt.Println(err)
	}
	Connection.CloseTCPConnection(c, pool)
	return message
}
