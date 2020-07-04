package Connection

// 通信インタフェース
type ConnectionInterface interface {
	AddAccessEvent(AccessEventInterface)		// アクセスイベント追加
	Service(string)		// 実行開始
}
