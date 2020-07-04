package InfrastructureConnection

import (
	"net/http"
	"github.com/gin-gonic/gin"
//	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/connection"
)

// LINEからの接続
type LINEAccessEvent struct {}

// コンストラクタ
func NewLINEAccessEvent() *LINEAccessEvent {
	var accessEvent = new(LINEAccessEvent)
	return accessEvent
}

// 情報取得
func (this *LINEAccessEvent) GetInfo() (Connection.AccessMethod, string) {
	return Connection.POST, "/line_client"
}

// 接続された時に走るイベント
func (this *LINEAccessEvent) OnAccess(gtx interface{}) {
	context := gtx.(*gin.Context)
	context.String(http.StatusOK, "LINE Access Event Running...")
}
