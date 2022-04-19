package model

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UserID            uint   `json:"user_id" form:"user_id" validate:"required"`
	Subscription_Date string `json:"subscription_date" form:"subscription_date"`
	Price             int    `json:"price" form:"price"`
	Description       string `json:"description" form:"description"`
}
