package connectionPool

import (
	"fmt"
	"net"
	"time"
)

var (
	Host                = "127.0.0.1"
	Port                = "8081"
	MIN_NUM_CONNECTIONS = 10
	MAX_NUM_CONNECTIONS = 100
)

func ConnectToTCPServer(pool *GncpPool) (net.Conn, error) {
	conn, err := pool.Get()
	if err != nil {
		fmt.Println("Establishing connection with timeout")
		conn, err = pool.GetWithTimeout(time.Duration(5) * time.Second)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return conn, nil
}

func CloseTCPConnection(conn net.Conn, pool *GncpPool) {
	err := conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pool.Remove(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ConnCreator() (net.Conn, error) {
	return net.Dial("tcp", Host+":"+Port)
}
