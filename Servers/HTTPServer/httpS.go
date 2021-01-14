package main

import (
	"log"
	"net/http"
)

func main() {
	// Registering the 2 new handler functions and corresponding URL
	// patterns with the servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/login", loginUser)
	mux.HandleFunc("/logout", logoutUser)
	mux.HandleFunc("/profile", showProfile)
	mux.HandleFunc("/profile/update", updateProfile)
	mux.HandleFunc("/uploadProfilePicture", uploadPicture)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	// CONNECT := ":8081"
	// c, err := net.Dial("tcp", CONNECT)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Connected...")
	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print(">> ")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Fprintf(c, text+"\n")
	// 	// fmt.Fprintf(c, "Hello..."+"\n")

	// 	//Receiving message from the TCP server
	// 	message, _ := bufio.NewReader(c).ReadString('\n')
	// 	fmt.Print("->: " + message)
	// }
}
