package entities

type User struct {
	name  string
	email string
}

func New(name string, email string) User {
	return User{name, email}
}
