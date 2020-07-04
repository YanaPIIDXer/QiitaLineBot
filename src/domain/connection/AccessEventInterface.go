package Connection

// メソッド
type AccessMethod int
const (
	POST AccessMethod = iota
	GET = iota
)

// アクセスイベントインタフェース
type AccessEventInterface interface {
	GetInfo() (AccessMethod, string)	// 情報取得
	OnAccess(interface{})				// 接続された時に走るイベント
}
