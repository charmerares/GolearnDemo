package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/hello/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello %s", name)
	})

	router.Run(":8080")
}
