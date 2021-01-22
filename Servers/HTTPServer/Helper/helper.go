package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

//TCP Communication functions
func ConnectToTCPServer() net.Conn {
	CONNECT := ":8081"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func CloseTCPConnection(conn net.Conn) {
	conn.Close()
}

func GetResponseFromTCPServer(command string, c net.Conn) string {
	//Text is the command to be sent to the TCP server
	text := command
	fmt.Fprintf(c, text+"\n")

	//Receiving message from the TCP server
	message, _ := bufio.NewReader(c).ReadString('\n')
	CloseTCPConnection(c)
	return message
}

func ConvertStringToMap(message string) (map[string]string, error) {
	details := make(map[string]string)
	err := json.Unmarshal([]byte(message), &details)
	if err != nil {
		return nil, err
	}
	return details, nil
}
