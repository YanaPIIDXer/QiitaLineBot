package Article

// 投稿データ
type Article struct {
	title *ArticleTitle		// タイトル
}

// コンストラクタ
func NewArticle(title *ArticleTitle) *Article {
	if title == nil { return nil }
	
	var article *Article = new(Article)
	article.title = title
	
	return article
}
