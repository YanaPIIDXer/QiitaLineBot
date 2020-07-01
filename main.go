package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

	// ローカルからの接続
    router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "HTTP SERVER")
	})
	
	// LINEからのリクエスト
	router.POST("/line_client", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Line Client")
	})

    router.Run(":3000")
}
