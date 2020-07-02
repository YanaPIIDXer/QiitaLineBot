package test

import (
	"testing"
	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/article"
)

// Article
func TestArticle(t *testing.T) {
	// インスタンス化
	var article = Article.NewArticle(Article.NewArticleTitle("Test"))
	if article == nil { t.Error("Article Instantiate Failed...") }

	// エラーチェック
	article = Article.NewArticle(nil)
	if article != nil { t.Error("Article Construct Validate Failed...") }

	t.Log("Article Test Success.")
}

// ArticleTitle
func TestArticleTitle(t *testing.T) {
	// インスタンス化
	var articleTitle = Article.NewArticleTitle("test")
	if articleTitle == nil { t.Error("ArticleTitle Instantiate Failed...") }

	// エラーチェック
	articleTitle = Article.NewArticleTitle("")
	if articleTitle != nil { t.Error("ArticleTitle Construct Validate Failed...") }

	// 等価チェック
	articleTitle = Article.NewArticleTitle("test")
	var sameTitle = Article.NewArticleTitle("test")
	if !articleTitle.Equal(sameTitle) { t.Error("ArticleTitle Equal Test Failed...") }
	var diffTitle = Article.NewArticleTitle("test2")
	if articleTitle.Equal(diffTitle) { t.Error("ArticleTitle Not Equal Test Failed...") }

	t.Log("ArticleTitle Test Success")
}
