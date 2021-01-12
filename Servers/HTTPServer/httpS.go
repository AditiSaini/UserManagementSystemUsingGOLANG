package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Home page"))
}

//Login handler function
func loginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logging in user..."))
}

//Logout handler function
func logoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logging out user..."))
}

//Show profile handler function
func showProfile(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Modify user profile..."))
}

//Upload profile picture
func uploadPicture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Uploading profile picture..."))
}

func main() {
	//Registering the 2 new handler functions and corresponding URL
	//patterns with the servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", loginUser)
	mux.HandleFunc("/logout", logoutUser)
	mux.HandleFunc("/profile", showProfile)
	mux.HandleFunc("/profile/update", updateProfile)
	mux.HandleFunc("/uploadProfilePicture", uploadPicture)

	log.Println("Staring server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
