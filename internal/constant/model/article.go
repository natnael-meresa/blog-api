package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Image  pq.StringArray `json:"image" form:"image"  gorm:"type:string[]"`
	Title  string         `json:"title" form:"title" validate:"required"`
	Tag    string         `json:"tag" form:"tag" validate:"required"`
	Status string         `json:"status" form:"status" validate:"required" gorm:"default:'pending'"`
	UserID uint           `json:"userid" form:"userid" validate:"required"`
}
