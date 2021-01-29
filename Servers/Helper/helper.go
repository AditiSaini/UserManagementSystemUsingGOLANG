package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"

	Connection "servers/ConnectionPool"
	Structure "servers/Structure"
)

func SendToHTTPServer(conn net.Conn, response string) {
	//Connecting to the HTTP Server to send the response
	conn.Write([]byte(response))
	conn.Close()
}

func ConvertStringToMap(message string) (map[string]string, error) {
	details := make(map[string]string)
	err := json.Unmarshal([]byte(message), &details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

//Hashes the password sent by the user
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func ConvertStructToMap(profile Structure.Profile) map[string]string {
	m := make(map[string]string)
	m["Username"] = profile.Username
	m["Nickname"] = profile.Nickname
	m["ImageRef"] = profile.ImageRef
	return m
}

func Visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
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
