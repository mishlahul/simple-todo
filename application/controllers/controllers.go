package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mishlahul/simple-todo/application/models"
)

type CreateUserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateTodoInput struct {
	Description string `json:"description" binding:"required"`
	IsCompleted bool   `json:"iscompleted"`
}

type UpdateUserInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UpdateTodoInput struct {
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

// GET /users
// get all users
func FindAllUser(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /todos
// get all todo items
func FindAllTodoItem(c *gin.Context) {
	var todoItem []models.TodoItem
	models.DB.Find(&todoItem)

	c.JSON(http.StatusOK, gin.H{"data": todoItem})
}

// POST /users
// create new user
func CreateUser(c *gin.Context) {
	// validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create user
	user := models.User{UserName: input.UserName, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /todos
// create new todo item
func CreateTodoItem(c *gin.Context) {
	// validate input
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create todo item
	todoItem := models.TodoItem{Description: input.Description, IsCompleted: input.IsCompleted}
	models.DB.Create(&todoItem)

	c.JSON(http.StatusOK, gin.H{"data": todoItem})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /todos/:id
// Find a todo item
func FindTodoItem(c *gin.Context) { // Get model if exist
	var todoItem models.TodoItem

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todoItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todoItem})
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /todos/:id
// Update a todo item
func UpdateTodoItem(c *gin.Context) {
	// Get model if exist
	var todoItem models.TodoItem
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todoItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todoItem).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todoItem})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DELETE /todos/:id
// Delete a todo item
func DeleteTodoItem(c *gin.Context) {
	// Get model if exist
	var todoItem models.TodoItem
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todoItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todoItem)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
