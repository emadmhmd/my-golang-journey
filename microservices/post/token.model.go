package main

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string  `json:"title" binding:"required"`
	Desc   string  `json:"desc" binding:"required"`
	UserId *string `json:"userId" binding:"required"`
}
