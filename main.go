package main

import (
	"go-library-management/controllers"
	"go-library-management/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.CreateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/list-users", controllers.ListUsers)
	router.Run(":8080")
}