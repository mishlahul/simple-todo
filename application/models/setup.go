package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := "host=localhost user=postgres password=postgres dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open("postgres", dbURL)

	if err != nil {
		panic("Failed to connect database!")
	}

	database.DropTableIfExists(&User{}, &TodoItem{})
	database.AutoMigrate(&User{}, &TodoItem{})

	DB = database
}
