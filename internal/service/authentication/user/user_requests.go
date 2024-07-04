package user

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=300"`
	Email    string `json:"email" validate:"required,email,max=100"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"-"`
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=300"`
	Email    string `json:"email" validate:"required,email,max=100"`
}
