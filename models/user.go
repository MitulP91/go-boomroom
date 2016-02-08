package "user"

type User struct {
	Username string
	Email string
}

// Constructor method
func NewUser() *User {
	return &User{}
}