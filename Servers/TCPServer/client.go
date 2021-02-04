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

	Helper "servers/TCPServer/Helper"
	Structure "servers/TCPServer/Structure"
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
		Helper.SendToHTTPServer(c.conn, "Send a recognizable command to the TCP Server")
	}
}

func (c *client) showUploadedFile(command *Structure.Command) {
	fmt.Println("Sending uploaded file")
	tokenAuth := command.Body["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	username, err := Helper.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
		return
	}

	//Get the filename with the username
	var files []string
	root := "./Pictures"
	err = filepath.Walk(root, Helper.Visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fileName := strings.Split(file, "/")[1]
		name := strings.Split(fileName, "-")[0]
		if name == username {
			fileName := file
			//Serve the image to the client
			//Read all the content of the uploaded file into a byte array
			byteFile, err := ioutil.ReadFile(fileName)
			//Convert the byte array into base64
			fileBase64 := base64.StdEncoding.EncodeToString([]byte(byteFile))
			if err != nil {
				fmt.Print(err)
			}
			Helper.SendToHTTPServer(c.conn, fileBase64)
			return
		}
	}
	Helper.SendToHTTPServer(c.conn, "false")
	return
}

func (c *client) receiveUploadedFile(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	encodedByteFile := command.Body["file"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	username, err := Helper.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
		return
	}

	//Delete all the other files with the same username
	var files []string
	root := "./Pictures"
	err = filepath.Walk(root, Helper.Visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fileName := strings.Split(file, "/")[1]
		name := strings.Split(fileName, "-")[0]
		if name == username {
			e := os.Remove(file)
			if e != nil {
				log.Fatal(e)
			}
		}
	}

	//Create a temp file
	name := username + "-*.png"
	tempFile, err := ioutil.TempFile("Pictures", name)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	decoded, err := base64.StdEncoding.DecodeString(encodedByteFile)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	//Write the byte array to the temp file
	tempFile.Write([]byte(decoded))

	Helper.SendToHTTPServer(c.conn, "true")
	return
}

func (c *client) changePassword(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	password := command.Body["password"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	username, err := Helper.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
		return
	}
	//Hash the password before storing into the database
	hashed, _ := Helper.HashPassword(password)
	updated, err := Helper.UpdatePassword(hashed, username)
	if !updated {
		Helper.SendToHTTPServer(c.conn, "false")
		return
	}
	Helper.SendToHTTPServer(c.conn, "true")
	return
}

func (c *client) updateProfile(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	name := command.Body["name"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	username, err := Helper.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
		return
	}
	updated, err := Helper.UpdateProfile(username, name)
	if !updated {
		Helper.SendToHTTPServer(c.conn, "false")
		return
	}
	Helper.SendToHTTPServer(c.conn, "true")
	return
}

func (c *client) logout(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)
	deleted, delErr := Helper.DeleteAuth(tokenAuthMap["AccessUUID"])
	if delErr != nil || deleted == 0 { //if anything goes wrong
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
		return
	}
	Helper.SendToHTTPServer(c.conn, "Logged out!")
	return
}

func (c *client) showProfile(command *Structure.Command) {
	tokenAuth := command.Body["tokenAuth"]
	tokenAuthMap, _ := Helper.ConvertStringToMap(tokenAuth)

	username, err := Helper.FetchAuth(tokenAuthMap)
	if err != nil {
		Helper.SendToHTTPServer(c.conn, "Unauthorised access")
		return
	}
	//Function to display the profile of the user from the database
	profile := Helper.Show(username)
	profileMap := Helper.ConvertStructToMap(profile)

	out, _ := json.Marshal(profileMap)
	Helper.SendToHTTPServer(c.conn, string(out))
	return
}

func (c *client) login(command *Structure.Command) {
	//Checking the validity of the credentials
	check := Helper.ValidateLogin(command.Body["username"], command.Body["password"])
	//If user is authenticated, get a bearer token and return it to the HTTP Server
	if check == true {
		//Token is created for auth
		token, _ := Helper.CreateToken(command.Body["username"])
		//User is saved in redis
		saveErr := Helper.CreateAuth(command.Body["username"], token)
		if saveErr != nil {
			tokens := map[string]string{
				"access_token": "Invalid Credentials",
			}
			out, _ := json.Marshal(tokens)
			Helper.SendToHTTPServer(c.conn, string(out))
			return
		}
		//Data prepared for sending to HTTP server
		tokens := map[string]string{
			"access_token": token.AccessToken,
		}
		out, _ := json.Marshal(tokens)
		//Data sent to HTTP server
		Helper.SendToHTTPServer(c.conn, string(out))
		return
	} else {
		tokens := map[string]string{
			"access_token": "Invalid Credentials",
		}
		out, _ := json.Marshal(tokens)
		Helper.SendToHTTPServer(c.conn, string(out))
		return
	}
}
