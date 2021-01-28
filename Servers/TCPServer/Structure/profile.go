package structure

type Profile struct {
	ID             int64
	Username       string
	Nickname       string
	Password       []byte
	ProfilePicture string
	Valid          bool
}
