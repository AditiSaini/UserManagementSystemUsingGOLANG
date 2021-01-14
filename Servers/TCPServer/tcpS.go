package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Printf("%v", err)
	}

	hub := newHub()
	go hub.run()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
		} else {
			fmt.Println("Connection Accepted...")
		}

		c := newClient(
			conn,
			hub.commands,
			"Adris",
		)
		go c.read()
	}
}

// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// 	"strconv"
// )

// var count = 0

// func handleConnection(c net.Conn) {
// 	fmt.Print(".")
// 	for {
// 		//Gets the data from the established connection
// 		// netData, err := bufio.NewReader(c).ReadString('\n')
// 		buf := make([]byte, 1024)
// 		netData, err := c.Read(buf)
// 		//Checks for errors
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		//Extracts the string enterd by the user
// 		// temp := strings.TrimSpace(string(netData))
// 		// if temp == "STOP" {
// 		// 	break
// 		// }
// 		//Prints the message
// 		fmt.Println(netData)
// 		//Increments the counter
// 		counter := strconv.Itoa(count) + "\n"
// 		//Sends is back to the client
// 		c.Write([]byte(string(counter)))
// 	}
// 	c.Close()
// }

// func main() {
// 	arguments := os.Args
// 	if len(arguments) == 1 {
// 		fmt.Println("Please provide port number")
// 		return
// 	}

// 	PORT := ":" + arguments[1]
// 	l, err := net.Listen("tcp4", PORT)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//Close the listener when application closes
// 	defer l.Close()

// 	for {
// 		c, err := l.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		go handleConnection(c)
// 		count++
// 	}
// }
