package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-todo-app/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	var todos []model.Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.BindJSON(&todo)
	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&model.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errrrrror": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
