package model

import (
	// "database/sql"
	// "time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name   string  `json:"name" binding:"required"`
	Email  *string `json:"email" binding:"required"`
	Mobile string  `json:"mobile"`
}
