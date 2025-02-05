package entities

type User struct {
	ID       int32
	Username string
	Email    string
}

func NewUser(username, email string) *User {
	return &User{ID: 1, Username: username, Email: email}
}
