package model

import (
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	ID string `json:"id" gorm:"primarykey"`
}
