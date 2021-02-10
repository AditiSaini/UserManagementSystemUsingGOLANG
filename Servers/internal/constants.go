package internal

const (
	// HTTP Server
	//1. Connection Pool (In both http and tcp server)
	HOST                = "127.0.0.1"
	TCP_PORT            = "8081"
	MIN_NUM_CONNECTIONS = 10
	MAX_NUM_CONNECTIONS = 5000
	NETWORK             = "tcp"

	//2. Helper (In both http and tcp server)
	TOKEN_SECRET = "secret"

	//3. httpS.go
	HTTP_PORT = "4000"

	//TCP Server
	//1. Helper
	DB_DRIVER = "mysql"
	DB_USER   = "root"
	DB_PASS   = ""
	DB_NAME   = "users"
	REDIS_DSN = "localhost:6379"
)
