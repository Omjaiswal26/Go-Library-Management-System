package controllers

import (
	"go-library-management/database"
	"go-library-management/models"
	"go-library-management/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	utils.Response(c, 200, true, "Books fetched successfully", books)
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err:= c.ShouldBindJSON(&input); err != nil {
		utils.Response(c, 400, false, "Invalid payload", nil)
		return
	}

	book := models.Book{
		Title: input.Title,
		Author: input.Author,
		Description: input.Description,
		Available: true,
	}

	if err := database.DB.Create(&book).Error; err != nil{
		log.Println("Error creating book", err)
		utils.InternalServerErrorResponse(c)
	}

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