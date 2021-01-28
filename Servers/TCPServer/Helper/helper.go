package helper

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"

	Structure "servers/TCPServer/Structure"
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
