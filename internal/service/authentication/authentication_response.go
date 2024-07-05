package authentication

type RegisterUserResponse struct {
	Email    string
	Username string
}

type LoginUserResponse struct {
	Username string
	Token    string
}
