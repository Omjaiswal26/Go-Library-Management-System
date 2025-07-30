package controllers

import (
	"go-library-management/database"
	"go-library-management/models"
	"go-library-management/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)


func IssueBook(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
		BookID uint `json:"book_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Response(c, 400, false, "BookID or UserID missing", nil)
		return
	}

	var book models.Book
	if err := database.DB.First(&book, input.BookID).Error; err != nil {
		utils.Response(c, 400, false, "No book found with given ID", nil)
		return
	}

	var user models.User
	if err := database.DB.First(&user, input.UserID).Error; err != nil {
		utils.Response(c, 400, false, "No user found with given ID", nil)
		return
	}

	if !book.Available {
		utils.Response(c, 400, false, "Book not available", nil)
		return
	}

	issue := models.BookIssue{
		UserID: input.UserID,
		BookID: input.BookID,
		IssueDate: time.Now(),
		ReturnDate: time.Now().AddDate(0, 0, 14),
	}

	if err := database.DB.Create(&issue).Error; err != nil {
		log.Print("Error issuing book", err)
		utils.Response(c, 500, false, "Error issuing book", nil)
		return
	}

	book.Available = false
	database.DB.Save(&book)
	database.DB.Preload("User").Preload("Book").First(&issue, issue.ID)

	utils.Response(c, 200, true, "Book issued successfully", &issue)
}