package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Description string
	IsCompleted bool
}
