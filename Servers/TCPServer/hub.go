package main

import (
	"fmt"
	"net"
)

type hub struct {
	commands chan command
}

func newHub() *hub {
	return &hub{
		commands: make(chan command),
	}
}

func (h *hub) run() {
	for {
		select {
		case cmd := <-h.commands:
			switch cmd.id {
			case nil:
				h.login(cmd.conn)
			default:
				fmt.Println("In hub default logic...")
				h.login(cmd.conn)
			}
		}
	}
}

func (h *hub) login(conn net.Conn) {
	fmt.Println("In login method...")
	conn.Write([]byte("Ok, logged in!"))
	conn.Close()
}
