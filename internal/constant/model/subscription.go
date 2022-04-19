package model

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserID            uint   `json:"user_id" form:"user_id" validate:"required"`
	ArticleID         uint   `json:"article_id" form:"article_id" validate:"required"`
	Subscription_Date string `json:"subscription_date" form:"subscription_date" gorm:"default:datatypes.Date(time.Now())"`
}
