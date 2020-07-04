package InfrastructureConnection

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/connection"
)

// LINEからの接続
type LINEAccessEvent struct {
	bot *linebot.Client		// BOTインスタンス
}

// コンストラクタ
func NewLINEAccessEvent() *LINEAccessEvent {
	var accessEvent = new(LINEAccessEvent)
	
	var channelSecret = os.Getenv("CHANNEL_SECRET")
	var channelToken = os.Getenv("CHANNEL_TOKEN")
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		fmt.Println("Line Bot Instantiate Failed.")
		fmt.Println("Channel Secret:" + channelSecret)
		fmt.Println("Channel Token:" + channelToken)
		return nil
	}
	accessEvent.bot = bot

	return accessEvent
}

// 情報取得
func (this *LINEAccessEvent) GetInfo() (Connection.AccessMethod, string) {
	return Connection.POST, "/line_client"
}

// 接続された時に走るイベント
func (this *LINEAccessEvent) OnAccess(gtx interface{}) {
	context := gtx.(*gin.Context)
	events, err := this.bot.ParseRequest(context.Request)
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
				_, err = this.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
				if err != nil {
					fmt.Println("ReplyMessage Failed.")
				}
				break
		}
	}

}
