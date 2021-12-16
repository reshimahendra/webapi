package moduls

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Book struct {
    gorm.Model
    ID          int             `json:"id"`
    Title       string          `json:"title" binding:"required"` 
    Description string          `json:"description"`
    Price       int             `json:"price" binding:"required,number,gte=0"`
    Rating      int             `json:"rating" binding:"required,number,gte=0,lte=5"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`
}

type BookRequest struct{
    Title        string         `json:"title" binding:"required"`
    Description  string         `json:"description"`
    Price        json.Number    `json:"price" binding:"required,number,gt=0"`
    Rating       json.Number    `json:"rating" binding:"required,number,gte=0,lte=5"`
}

type BookResponse struct{
    ID           int            `json:"id"`
    Title        string         `json:"title"`
    Description  string         `json:"description"`
    Price        int            `json:"price"`
    Rating       int            `json:"rating"`
}

