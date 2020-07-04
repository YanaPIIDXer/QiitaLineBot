package InfrastructureConnection

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/connection"
)

// ローカルからの接続
type LocalAccessEvent struct {}

// コンストラクタ
func NewLocalAccessEvent() *LocalAccessEvent {
	var accessEvent = new(LocalAccessEvent)
	return accessEvent
}

// 情報取得
func (this *LocalAccessEvent) GetInfo() (Connection.AccessMethod, string) {
	return Connection.GET, "/"
}

// 接続された時に走るイベント
func (this *LocalAccessEvent) OnAccess(gtx interface{}) {
	context := gtx.(*gin.Context)
	context.String(http.StatusOK, "Local Access Event Running...")
}
