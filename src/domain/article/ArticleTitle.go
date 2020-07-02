package Article

// 投稿タイトル
type ArticleTitle struct {
	value string		// 値
}

// コンストラクタ
func NewArticleTitle(value string) *ArticleTitle {
	if value == "" { return nil }
	
	var articleTitle *ArticleTitle = new(ArticleTitle)
	articleTitle.value = value
	
	return articleTitle
}

// 値を取得
func (this *ArticleTitle) Value() string {
	return this.value;
}

// 等価判定
func (this *ArticleTitle) Equal(other *ArticleTitle) bool {
	return (this.value == other.value)
}
