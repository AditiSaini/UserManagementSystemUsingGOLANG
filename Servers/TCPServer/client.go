package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	Auth "servers/Authentication"
	Helper "servers/Helper"
	MySQL "servers/MySQL"
	Redis "servers/Redis"
	Structure "servers/Structure"
)

var (
	DELIMITER   = []byte(`\r\n`)
	imageUpload = false
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
			fmt.Println(err)
			return err
		}
		c.handle(msg)
	}
}

func (c *client) handle(message []byte) {
	//Processing commands sent by the HTTP Server
	//Step 1- Get the command
	cmd := bytes.ToUpper(bytes.TrimSpace(bytes.Split(message, []byte(" "))[0]))
	//Step 2- Extract the arguments of the command
	args := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))
	//Step 3- Processing the arguments of the command
	processArgs := Helper.ExtractingArgumentsFromCommands(string(cmd), string(args))
	//Converted into the command data structure
	command := Structure.NewCmd(string(cmd), processArgs, c.conn)

	fmt.Println("Handling command: " + string(cmd))

	//Routing the command to the right handler function
	switch string(cmd) {
	case "LOGIN":
		c.login(command)
	case "SHOW_PROFILE":
		c.showProfile(command)
	case "LOGOUT":
		c.logout(command)
	case "UPDATE_PROFILE":
		c.updateProfile(command)
	case "CHANGE_PASSWORD":
		c.changePassword(command)
	case "UPLOAD_PICTURE":
		c.receiveUploadedFile(command)
	case "SHOW_PICTURE":
		c.showUploadedFile(command)
	default:
		Helper.SendToHTTPServer(c.conn, "Send a recognizable command to the TCP Server\n")
	}
}

func (c *client) showUploadedFile(command *Structure.Command) {
	fmt.Println("Sending uploaded file")
	tokenAuth := command.Body["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	profile, err := Redis.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}
	//Function to display the profile of the user from the database
	// profile := MySQL.Show(username)
	profileMap := Helper.ConvertStructToMap(profile)
	//Read all the content of the uploaded file into a byte array
	byteFile, err := ioutil.ReadFile(profileMap["ImageRef"])
	//Convert the byte array into base64
	fileBase64 := base64.StdEncoding.EncodeToString([]byte(byteFile))
	if err != nil {
		fmt.Print(err)
		Helper.SendToHTTPServer(c.conn, "false\n")
		return
	}
	Helper.SendToHTTPServer(c.conn, fileBase64+"\n")
	return
}

func (c *client) receiveUploadedFile(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	encodedByteFile := command.Body["file"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	profile, err := Redis.FetchAuth(tokenAuthMap)
	username := profile.Nickname
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}

	//Delete all the other files with the same username
	var files []string
	root := "../Pictures"
	err = filepath.Walk(root, Helper.Visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		list := strings.Split(file, "/")
		if len(list) > 2 {
			fileName := strings.Split(file, "/")[2]
			name := strings.Split(fileName, "-")[0]
			if name == username {
				e := os.Remove(file)
				if e != nil {
					log.Fatal(e)
				}
			}
		}
	}

	//Create a temp file
	name := username + "-*.png"
	tempFile, err := ioutil.TempFile("../Pictures", name)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	finalName := tempFile.Name()
	profile.ImageRef = finalName
	//Updates user in redis
	saveErr := Redis.UpdateUserProfile(profile, tokenAuthMap)
	if saveErr != nil {
		Helper.SendToHTTPServer(c.conn, "false\n")
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(encodedByteFile)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	//Write the byte array to the temp file
	tempFile.Write([]byte(decoded))

	Helper.SendToHTTPServer(c.conn, "true\n")
	return
}

func (c *client) changePassword(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	password := command.Body["password"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	profile, err := Redis.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}

	//Hash the password before storing into the database
	hashed, _ := Helper.HashPassword(password)
	profile.Password = hashed
	//Updates user in redis
	saveErr := Redis.UpdateUserProfile(profile, tokenAuthMap)
	if saveErr != nil {
		Helper.SendToHTTPServer(c.conn, "false\n")
		return
	}
	Helper.SendToHTTPServer(c.conn, "true\n")
	return
}

func (c *client) updateProfile(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	name := command.Body["name"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	profile, err := Redis.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}
	profile.Username = name
	//Updates user in redis
	saveErr := Redis.UpdateUserProfile(profile, tokenAuthMap)
	if saveErr != nil {
		Helper.SendToHTTPServer(c.conn, "false\n")
		return
	}
	Helper.SendToHTTPServer(c.conn, "true\n")
	return
}

//Update the data in the db
func (c *client) logout(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	//Fetch user profile from redis
	profile, err := Redis.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}
	//Update this user profile in Mysql db
	updated, err := MySQL.UpdateUserProfile(profile)
	if !updated {
		Helper.SendToHTTPServer(c.conn, "false\n")
		return
	}

	deleted, delErr := Redis.DeleteAuth(tokenAuthMap["AccessUUID"])
	if delErr != nil || deleted == 0 { //if anything goes wrong
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}
	Helper.SendToHTTPServer(c.conn, "true\n")
	return
}

func (c *client) showProfile(command *Structure.Command) {
	fmt.Println(command.Body["tokenAuth"])
	tokenAuth := command.Body["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)

	profile, err := Redis.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access\n")
		return
	}
	//Function to display the profile of the user from the database
	profileMap := Helper.ConvertStructToMap(profile)
	out, _ := json.Marshal(profileMap)
	Helper.SendToHTTPServer(c.conn, string(out)+"\n")
	return
}

func (c *client) login(command *Structure.Command) {
	//Checking the validity of the credentials
	profile, err := Auth.ValidateLogin(command.Body["username"], command.Body["password"])
	//If user is authenticated, get a bearer token and return it to the HTTP Server
	if err == nil {
		//Token is created for auth
		token, _ := Auth.CreateToken(command.Body["username"])
		//User is saved in redis
		saveErr := Redis.CreateAuth(profile, token)
		if saveErr != nil {
			tokens := map[string]string{
				"access_token": "Invalid Credentials",
			}
			out, _ := json.Marshal(tokens)
			Helper.SendToHTTPServer(c.conn, string(out)+"\n")
			return
		}
		//Data prepared for sending to HTTP server
		tokens := map[string]string{
			"access_token": token.AccessToken,
		}
		out, _ := json.Marshal(tokens)
		//Data sent to HTTP server
		Helper.SendToHTTPServer(c.conn, string(out)+"\n")
		return
	} else {
		tokens := map[string]string{
			"access_token": "Invalid Credentials",
		}
		out, _ := json.Marshal(tokens)
		Helper.SendToHTTPServer(c.conn, string(out)+"\n")
		return
	}
}
