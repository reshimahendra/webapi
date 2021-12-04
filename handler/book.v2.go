package handler

import (
	"fmt"
	"net/http"
    "webapi/entry"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BookPostHandler2(c *gin.Context) {
    var book entry.BookInput
    err := c.ShouldBindJSON(&book)

    if err != nil {
        eMsg := []string{} 
        for _,e := range err.(validator.ValidationErrors) {
            msg := fmt.Sprintf("Error on field '%s', condition: '%s'", e.Field(), e.ActualTag)
            eMsg = append(eMsg, msg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "title": book.Title,
        "sub_title": book.SubTitle,
        "price": book.Price,
        "description": book.Description,
        "release_date": book.Release_date,
    })
}
