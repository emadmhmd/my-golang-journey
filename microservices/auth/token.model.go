package main

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token  string  `json:"token" binding:"required"`
	UserId *string `json:"userId" binding:"required"`
}
