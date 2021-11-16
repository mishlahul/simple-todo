package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func DbInit() {
	log.Info("Starting Database Connection")
	dbURL := "host=localhost user=mmuzakki password=postgres dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Panic("Failed to connect database with error", err.Error())
	}

	db.LogMode(true)
}
