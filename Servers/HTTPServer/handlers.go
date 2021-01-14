package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func connectToTCPServer() net.Conn {
	CONNECT := ":8081"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func getResponseFromTCPServer(c net.Conn) string {
	// for {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print(">> ")
	// text, _ := reader.ReadString('\n')
	text := "Hello "
	fmt.Fprintf(c, text+"\n")
	// fmt.Fprintf(c, "Hello...")

	//Receiving message from the TCP server
	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Print("->: " + message)
	// if strings.TrimSpace(string(text)) == "STOP" {
	// 	fmt.Println("TCP client exiting...")
	// 	return
	// }
	c.Close()
	return message
	// }
}

func home(w http.ResponseWriter, r *http.Request) {
	c := connectToTCPServer()
	message := getResponseFromTCPServer(c)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Home page " + message))
}

//Login handler function
func loginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Logging in user..."))
}

//Logout handler function
func logoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logging out user..."))
}

//Show profile handler function
func showProfile(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Displaying profile...")
	w.Write([]byte("Displaying profile..."))
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
	w.Write([]byte("Modify user profile..."))
}

//Upload profile picture
func uploadPicture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Uploading profile picture..."))
}
