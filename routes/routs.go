package routs

import (
	"fmt"
	"go-todo-app/controller"
	"go-todo-app/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// PostgreSQL DSN
	dsn := "host=localhost user=sinan password=sinan123 dbname=todo_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Migrate and connect
	model.MigrateTodos(db)
	controller.SetDB(db)

	r.POST("/todos", controller.CreateTodo)
	r.GET("/todos", controller.GetTodos)
	r.PUT("/todos/:id", controller.UpdateTodo)
	r.DELETE("/todos/:id", controller.DeleteTodo)

	return r
}
