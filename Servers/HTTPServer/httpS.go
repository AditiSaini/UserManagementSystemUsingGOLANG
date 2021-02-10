package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Constants "servers/internal"
)

func main() {
	// Registering the 2 new handler functions and corresponding URL
	// patterns with the servemux
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/login", loginUser)
	router.HandleFunc("/logout", logoutUser)
	router.HandleFunc("/profile", showProfile)
	router.HandleFunc("/profile/update", updateProfile)
	router.HandleFunc("/uploadProfilePicture", uploadPicture)
	router.HandleFunc("/showProfilePicture", showPicture)
	router.HandleFunc("/changePassword", changePassword)

	log.Println("Starting server on :" + Constants.HTTP_PORT)
	err := http.ListenAndServe(":"+Constants.HTTP_PORT, router)
	log.Fatal(err)
}
