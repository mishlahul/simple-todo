package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TodoItem struct {
	gorm.Model
	Description string `json:"description"`
	IsCompleted bool   `json:"iscompleted"`
}
