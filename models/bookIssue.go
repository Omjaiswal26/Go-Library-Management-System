package models

import (
	"time"

	"gorm.io/gorm"
)

type BookIssue struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BookID uint `json:"book_id"`
	Book Book `json:"book" gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	IssueDate time.Time `json:"issue_date"`
	ReturnDate time.Time `json:"return_date"`
}