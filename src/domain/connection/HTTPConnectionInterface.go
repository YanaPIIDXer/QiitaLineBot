package Connection

// HTTP通信インタフェース
type HTTPConnectionInterface interface {
	AddAccessEvent(*AccessEventInterface)		// アクセスイベント追加
	Service(string)		// 実行開始
}
