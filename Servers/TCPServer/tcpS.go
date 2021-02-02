package main

import (
	"fmt"
	"log"
	"net"

	Constants "servers/internal"
)

func main() {
	ln, err := net.Listen(Constants.NETWORK, ":"+Constants.TCP_PORT)
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
