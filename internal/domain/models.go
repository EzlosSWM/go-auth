package domain

type User struct {
	ID uint `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password []byte `json:"password"`
}

func NewUser(first, last, email string, password []byte) User {
	return User{
		FirstName: first,
		LastName: last,
		Email: email,
		Password: password,
	}
}
