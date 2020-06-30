package Post

// 投稿タイトル
type PostTitle struct {
	value string		// 値
}

// コンストラクタ
func NewPostTitle(value string) *PostTitle {
	if value == "" { return nil }
	
	var postTitle *PostTitle = new(PostTitle)
	postTitle.value = value
	
	return postTitle
}

// 値を取得
func (this *PostTitle) Value() string {
	return this.value;
}

// 等価判定
func (this *PostTitle) Equal(other *PostTitle) bool {
	return (this.value == other.value)
}
