package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Registering the 2 new handler functions and corresponding URL
	// patterns with the servemux
	router := mux.NewRouter()
	// mux := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/login", loginUser)
	router.HandleFunc("/logout", logoutUser).Methods("GET")
	router.HandleFunc("/profile", showProfile).Methods("POST")
	router.HandleFunc("/profile/update", updateProfile).Methods("POST")
	router.HandleFunc("/uploadProfilePicture", uploadPicture).Methods("POST")

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}
