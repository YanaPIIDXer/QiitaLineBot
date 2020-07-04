package main

import (
	"os"
	"github.com/joho/godotenv"

	"github.com/YanaPIIDXer/QiitaLineBot/src/infrastructure/connection"
)

func main() {
	godotenv.Load(".env")

	var port = os.Getenv("PORT")		// heroku上で動かす場合は指定されたPortじゃないと死なますよ。
	if port == "" {
		port = "3000"
	}

	var connection = InfrastructureConnection.NewHTTPConnection()
	connection.AddAccessEvent(InfrastructureConnection.NewLocalAccessEvent())
	connection.AddAccessEvent(InfrastructureConnection.NewLINEAccessEvent())
	connection.Service(port)
	
	/*
	// この辺の処理、mainに直接書くのもなぁ・・・
	// @TODO:後程クラス化。
	router := gin.Default()
	
	var channelSecret = os.Getenv("CHANNEL_SECRET")
	var channelToken = os.Getenv("CHANNEL_TOKEN")
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		fmt.Println("Line Bot Instantiate Failed.")
		fmt.Println("Channel Secret:" + channelSecret)
		fmt.Println("Channel Token:" + channelToken)
		return
	}

	// ローカルからの接続
    router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "LOCAL SERVER")
	})
	
	// LINEからのリクエスト
	router.POST("/line_client", func(ctx *gin.Context) {
		events, err := bot.ParseRequest(ctx.Request)
		if err != nil {
			fmt.Println("Parse Request Failed.")
			if err == linebot.ErrInvalidSignature {
				fmt.Println(" ...Invalid Signature.")
			}
			return
		}

		for _, event := range events {
			switch message := event.Message.(type) {
				case *linebot.TextMessage:
					// ↓ReplyMessageメソッドが返すのはエラーオブジェクトではなくReplyMessageCallと言うオブジェクトで、
					//  ここでメッセージが送信されるわけではない。
					//  そいつのDoメソッドを叩くことで初めてメッセージが送信される。
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
					if err != nil {
						fmt.Println("ReplyMessage Failed.")
					}
					break
			}
		}
	})

	var port = os.Getenv("PORT")		// heroku上で動かす場合は指定されたPortじゃないと死なますよ。
	if port == "" {
		port = "3000"
	}
	fmt.Println("Server Port:" + port)
	router.Run(":" + port)
	*/
}
