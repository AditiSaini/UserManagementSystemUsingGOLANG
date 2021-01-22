package helper

import (
	"encoding/json"
	"net"

	"golang.org/x/crypto/bcrypt"
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
