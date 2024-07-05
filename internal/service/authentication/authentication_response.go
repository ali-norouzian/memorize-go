package authentication

type RegisterUserResponse struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Username string `json:"username" validate:"required,max=100"`
}
