package main

import (
    hdl "webapi/handler"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.SetTrustedProxies([]string{"192.168.43.26"})

    // membuat api versioning dengan nama v1 (api version 1)
    v1 := router.Group("/v1")

    v1.GET("/hello", hdl.HelloHandler) 
    v1.GET("/book/:id/:title", hdl.BookDetailHandler)
    v1.POST("/books", hdl.BookPostHandler)
    v1.GET("/search", hdl.SearchQueryHandler)
    v1.GET("/", hdl.RootHandler)

    v2 := router.Group("/v2")
    v2.POST("/books", hdl.BookPostHandler2)

    router.Run(":3888")
}

