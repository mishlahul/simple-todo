package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mishlahul/simple-todo/application/db"
	"github.com/mishlahul/simple-todo/application/handlers"
	"github.com/mishlahul/simple-todo/application/models"
	log "github.com/sirupsen/logrus"
	// "gorm.io/gorm"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todo API")

	db := db.DbInit()
	defer db.Close()
	db.DropTableIfExist(&models.TodoItem{}, &models.User{})
	db.AutoMigrate(&models.TodoItem{}, &models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", healthcheckHandler).Methods("GET")

	router.HandleFunc("/register", handlers.CreateUserHandler(db)).Methods("POST")
	router.HandleFunc("/users", handlers.GetListUserHandler(db)).Methods("GET")

	http.ListenAndServe(":8080", router)

}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
