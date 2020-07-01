package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "HTTP SERVER")
    })

    router.Run(":3000")
}
