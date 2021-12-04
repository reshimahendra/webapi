package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
    router := gin.Default()
    router.SetTrustedProxies([]string{"192.168.43.26"})

    router.GET("/hello", HelloHandler) 
    router.GET("/book/:id/:title", BookDetailHandler)
    router.POST("/books", BookPostHandler)
    router.GET("/search", SearchQueryHandler)
    router.GET("/", RootHandler)

    router.Run(":3888")
}

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

type BookInput struct{
    Title string `json:"title" binding:"required"`
    Price json.Number `json:"price" binding:"required,number,gt=0"`
    SubTitle string `json:"sub_title" binding:"required"`
}

func BookPostHandler(c *gin.Context) {
    var bookInput BookInput
    err := c.ShouldBindJSON(&bookInput)
    if err != nil {
        for _, e := range err.(validator.ValidationErrors) {
            eMsg := fmt.Sprintf("Error on field '%s', condition: '%s'.", e.Field(), e.ActualTag())
            c.JSON(http.StatusBadRequest, eMsg)
        }
        return
    }   

    c.JSON(http.StatusOK, gin.H{
        "title": bookInput.Title,
        "price": bookInput.Price,
        "sub_title": bookInput.SubTitle,
    })
}
