package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	Constants "servers/internal"
)

func main() {
	ln, err := net.Listen(Constants.NETWORK, ":"+Constants.TCP_PORT)
	if err != nil {
		log.Printf("%v", err)
	}
	defer ln.Close()
	rand.Seed(time.Now().Unix())

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
