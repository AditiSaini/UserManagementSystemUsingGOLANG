package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	Helper "./Helper"
	Structure "./Structure"
)

//Routing handler functions
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Connecting to the tcp server
	c := Helper.ConnectToTCPServer()
	message := Helper.GetResponseFromTCPServer("home page", c)
	//Sending information to the client
	w.Write([]byte("Hello from Home page!!!\n" + message))
}

//Login handler function
func loginUser(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	var user Structure.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	username := user.Username
	password := user.Password

	command := "LOGIN username " + username + "|password " + password
	c := Helper.ConnectToTCPServer()
	message := Helper.GetResponseFromTCPServer(command, c)
	details, _ := Helper.ConvertStringToMap(message)
	m["command"] = "LOGIN"
	m["access_token"] = details["access_token"]
	jsonString, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(jsonString))
}

//Logout handler function
func logoutUser(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	c := Helper.ConnectToTCPServer()
	tokenAuth, err := Helper.ExtractTokenMetadata(r)
	if err != nil {
		fmt.Println(err)
		m["profile"] = "Unauthorised Access"
		jsonString, _ := json.Marshal(m)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(jsonString))
		return
	}
	b, err := json.Marshal(tokenAuth)
	if err != nil {
		fmt.Println(err)
		return
	}
	command := "LOGOUT tokenAuth " + string(b)
	message := Helper.GetResponseFromTCPServer(command, c)
	m["command"] = "LOGOUT"
	m["status"] = message
	jsonString, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(jsonString))
}

//Show profile handler function
func showProfile(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	c := Helper.ConnectToTCPServer()
	tokenAuth, err := Helper.ExtractTokenMetadata(r)
	if err != nil {
		fmt.Println(err)
		m["profile"] = "Unauthorised Access"
		jsonString, _ := json.Marshal(m)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(jsonString))
		return
	}
	b, err := json.Marshal(tokenAuth)
	if err != nil {
		fmt.Println(err)
		return
	}
	command := "SHOW_PROFILE tokenAuth " + string(b)
	message := Helper.GetResponseFromTCPServer(command, c)
	details, _ := Helper.ConvertStringToMap(message)
	jsonString, _ := json.Marshal(details)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(jsonString)
}

//Modify profile handler function
func updateProfile(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note that
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	c := Helper.ConnectToTCPServer()
	message := Helper.GetResponseFromTCPServer("update profile handler method", c)
	w.Write([]byte("Modify user profile..." + message))
}

//Upload profile picture
func uploadPicture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	c := Helper.ConnectToTCPServer()
	message := Helper.GetResponseFromTCPServer("upload profile picture handler method", c)
	w.Write([]byte("Uploading profile picture..." + message))
}
