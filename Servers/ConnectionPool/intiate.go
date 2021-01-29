package connectionPool

import (
	"fmt"
	"net"
	"time"

	Constants "servers/internal"
)

var (
	Network             = Constants.NETWORK
	Host                = Constants.HOST
	Port                = Constants.TCP_PORT
	MIN_NUM_CONNECTIONS = Constants.MIN_NUM_CONNECTIONS
	MAX_NUM_CONNECTIONS = Constants.MAX_NUM_CONNECTIONS
)

func ConnectToTCPServer(pool *GncpPool) (net.Conn, error) {
	conn, err := pool.Get()
	if err != nil {
		fmt.Println("Establishing connection with timeout")
		conn, err = pool.GetWithTimeout(time.Duration(1) * time.Second)
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
		err = pool.Remove(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func ConnCreator() (net.Conn, error) {
	return net.Dial(Network, Host+":"+Port)
}

func InitialisePoolValue(pool *GncpPool) (*GncpPool, error) {
	if pool == nil {
		fmt.Println("Initialised pool")
		return NewPool(MIN_NUM_CONNECTIONS, MAX_NUM_CONNECTIONS, ConnCreator)
	}
	return pool, nil
}
