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

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
		} else {
			fmt.Println("Connection Accepted...")
		}
		c := newClient(
			conn,
		)
		go c.read()
	}
}
