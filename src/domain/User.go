package domain

type User struct {
	Username string
	Password string
	Tweets   []Tweet
}

func NewUser(username string, password string) *User {
	return &User{username, password, nil}
}
