package entities

type user struct {
	id    int32
	name  string
	email string
}

func NewUser(name, email string) *user {
	return &user{
		name:  name,
		email: email,
	}
}

func (u *user) GetID() int32 {
	return u.id
}

func (u *user) SetID(id int32) {
	u.id = id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) SetEmail(email string) {
	u.email = email
}
