package user

import "time"

type CreateUserResponse struct {
	ID uint
}

type ListUsersResponse struct {
	ID        uint
	Username  string
	Email     string
	CreatedAt time.Time
}

type GetUserByIDResponse struct {
	ID        uint
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
