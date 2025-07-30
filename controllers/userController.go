package controllers

import (
	"go-library-management/database"
	"go-library-management/models"
	"go-library-management/utils"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Response(c, 400, false, "Invalid Json", nil)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.Response(c, 400, false, "User already exists", nil)
		return
	}

	utils.Response(c, 200, true, "Registered Successfully", nil)
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Response(c, 400, false, "Invalid payload", nil)
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.Response(c, 400, false, "Invalid email or password", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.Response(c, 400, false, "Invalid email or password", nil)
		return
	}

	utils.Response(c, 200, true, "Login Successful", &user)
}

func ListUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		log.Println("Error fetching users: ", err)
		utils.Response(c, 500, false, "Internal Server Error", nil)
		return
	}

	utils.Response(c, 200, true, "Users fetched successfully", &users)
}


func UserIssuedBooks(c *gin.Context) {
	var user models.User
	var books []models.Book
	var issuedBooks []models.BookIssue

	id := c.Param("id")
	if err := database.DB.First(&user, id).Error; err != nil {
		utils.NotFoundResponse(c)
		return
	}

	if err := database.DB.Where("user_id = ?", user.ID).Preload("Book").Find(&issuedBooks).Error; err != nil {
		utils.NotFoundResponse(c)
		return
	}

	for _, issue := range issuedBooks {
		books = append(books, issue.Book)
	}

	utils.SuccessResponse(c, "Issued books for User fetched Successfully", books)

}
