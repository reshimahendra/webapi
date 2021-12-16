package handler

import (
	"fmt"
	"net/http"
    "webapi/moduls"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
func RootHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "name": "Reshi Mahendra",
        "Bio": "Spearfisher Man",
    })
}


func HelloHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "name": "Reshi",
        "bio": "An Internet marketter",
        "content": "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
    })
}

func BookDetailHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "id": c.Param("id"),
        "name" : c.Param("name"),
        "title": c.Param("title"),
    })
}

func SearchQueryHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "q": c.Query("q"),
        "price": c.Query("price"),
    })
}

func BookPostHandler(c *gin.Context) {
    var bookRequest moduls.BookRequest
    err := c.ShouldBindJSON(&bookRequest)

    eMsg := []string{}
    if err != nil {
        for _, e := range err.(validator.ValidationErrors) {
            msg := fmt.Sprintf("Error on field '%s', condition: '%s'.", e.Field(), e.ActualTag())
            eMsg = append(eMsg, msg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }   

    c.JSON(http.StatusOK, gin.H{
        "title": bookRequest.Title,
        "price": bookRequest.Price,
        "description": bookRequest.Description,
        "rating": bookRequest.Rating,
    })
}
