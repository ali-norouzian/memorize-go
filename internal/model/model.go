package model

func NewListOfDbModels() []any {
	return []any{
		&User{},
		&Question{},
		&QuestionUser{},
	}
}
