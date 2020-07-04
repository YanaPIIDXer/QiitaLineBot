package Connection

// アクセスイベントインタフェース
type AccessEventInterface interface {
	GetPath() (string)			// パス取得
	OnAccess(*interface{})		// 接続された時に走るイベント
}
