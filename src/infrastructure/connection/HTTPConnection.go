package Connection

import (
	"github.com/gin-gonic/gin"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/connection"
)

// HTTP通信
type HTTPConnection struct {
	router *gin.Engine
}

// コンストラクタ
func NewHTTPConnection() *HTTPConnection {
	var httpConnection = new(HTTPConnection)
	httpConnection.router = gin.Default()
	return httpConnection
}

// アクセスイベント追加
func (this *HTTPConnection) AddAccessEvent(event Connection.AccessEventInterface) {
	method, path := event.GetInfo()
	switch(method) {
		case Connection.POST:
			this.router.POST(path, func(ctx *gin.Context) { event.OnAccess(ctx) })
			break
		case Connection.GET:
			this.router.GET(path, func(ctx *gin.Context) { event.OnAccess(ctx) })
			break
	}
}

// 実行
func (this *HTTPConnection) Service(port string) {
	router.Run(":" + port)
}
