package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"name" form:"name" validate:"required"`
	Email    string     `json:"email" form:"email" validate:"required,email"`
	Password string     `json:"password" form:"password" validate:"required,gte=5,lte=12"`
	Role     string     `json:"role" form:"role" gorm:"-"`
	Article  *[]Article `validate:"omitempty"`
}
