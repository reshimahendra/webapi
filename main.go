package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.SetTrustedProxies([]string{"192.168.43.26"})

    router.GET("/hello", HelloHandler) 
    router.GET("/book/:id/:title", BookDetailHandler)
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
        "id" : c.Param("id"),
        "title": c.Param("title"),
    })
}

func SearchQueryHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "q": c.Query("q"),
        "price": c.Query("price"),
    })
}
