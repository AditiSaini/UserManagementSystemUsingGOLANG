package helper

import "net"

func SendToHTTPServer(conn net.Conn, response string) {
	//Connecting to the HTTP Server to send the response
	conn.Write([]byte(response))
	conn.Close()
}
