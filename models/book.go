package models

type Book struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
	Available bool `json:"available"`
}

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description"`
}
