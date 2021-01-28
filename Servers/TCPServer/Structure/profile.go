package structure

type Profile struct {
	ID       int64
	Username string
	Nickname string
	Password []byte
	Valid    bool
}
