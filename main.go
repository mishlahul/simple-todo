package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mishlahul/simple-todo/application/controllers"
	"github.com/mishlahul/simple-todo/application/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "todo-list"})
	})

	r.GET("/users", controllers.FindAllUser)
	r.GET("/todos", controllers.FindAllTodoItem)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/todos/:id", controllers.FindTodoItem)
	r.POST("/users", controllers.CreateUser)
	r.POST("/todos", controllers.CreateTodoItem)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.PATCH("/todos/:id", controllers.UpdateTodoItem)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.DELETE("/todos/:id", controllers.DeleteTodoItem)

	r.Run()

}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
