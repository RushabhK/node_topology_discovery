package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "server is live..")
	})

	r.Run(":8080")
}
