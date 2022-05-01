package database

type DatabasePaths struct {
	Root     string
	Etc      string
	Chat     string
	Users    string
	Channels string
	Key      string
}

type User struct {
	Name         string
	PasswordHash string
}

type Channel struct {
	Name    string
	Members []*User
}

var (
	Paths    DatabasePaths
	Users    []User
	channels []Channel
)
