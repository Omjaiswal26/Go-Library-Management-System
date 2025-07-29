package controllers

import (
	"go-library-management/database"
	"go-library-management/models"
	"go-library-management/utils"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	utils.Response(c, 200, true, "Books fetched successfully", books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err:= c.ShouldBindJSON(&book); err != nil {
		utils.Response(c, 400, false, "Invalid payload", nil)
		return
	}
	database.DB.Create(&book)
	utils.Response(c, 201, true, "Book created successfully", book)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err:= database.DB.First(&book, id).Error; err != nil {
		utils.Response(c, 400, false, "Book with given ID not found", nil)
		return
	}

	utils.Response(c, 200, true, "Book fetched successfully", book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err:= database.DB.First(&book, id).Error; err != nil {
		utils.Response(c, 400, false, "Book with given ID not found", nil)
		return
	}

	database.DB.Delete(&book)
	utils.Response(c, 200, true, "Book deleted successfully", nil)
}