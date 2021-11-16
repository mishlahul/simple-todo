package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/mishlahul/simple-todo/application/models"
)

func CreateUserHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		password := r.FormValue("password")

		newUser := &models.User{Name: name, Password: password}
		db.Create((newUser))
		result := db.Last(&newUser)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result.Value)
	}
	return http.HandlerFunc(fn)
}

func GetListUserHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		users := []models.User{}

		db.Find(&users)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
	return http.HandlerFunc(fn)
}
