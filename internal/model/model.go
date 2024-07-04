package model

import "memorize/internal/model/authentication"

func NewListOfDbModels() []any {
	return []any{
		&authentication.User{},
	}
}
