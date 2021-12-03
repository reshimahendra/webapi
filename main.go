package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.SetTrustedProxies([]string{"192.168.43.26"})

    router.GET("/hello", func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
            "name": "Reshi",
            "bio": "An Internet marketter",
            "content": "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
        })
    })

    router.GET("/", func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
            "name": "Reshi Mahendra",
            "Bio": "Spearfisher Man",
        })
    })

    router.Run(":3888")
}
