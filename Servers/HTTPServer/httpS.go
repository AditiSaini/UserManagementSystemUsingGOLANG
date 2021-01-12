package main

import (
	"log"
	"net/http"
)

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
