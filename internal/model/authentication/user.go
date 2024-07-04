package authentication

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"size:300;not null"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	// Questions []Question `gorm:"many2many:question_users;"`
}
