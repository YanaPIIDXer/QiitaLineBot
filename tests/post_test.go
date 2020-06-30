package test

import (
	"testing"
	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/post"
)

// Post
func TestPost(t *testing.T) {
	// インスタンス化
	var post = Post.NewPost(Post.NewPostTitle("Test"))
	if post == nil { t.Error("Post Instantiate Failed...") }

	// エラーチェック
	post = Post.NewPost(nil)
	if post != nil { t.Error("Post Construct Validate Failed...") }

	t.Log("Post Test Success.")
}

// PostTitle
func TestPostTitle(t *testing.T){
	// インスタンス化
	var postTitle = Post.NewPostTitle("test")
	if postTitle == nil { t.Error("PostTitle Instantiate Failed...") }

	// エラーチェック
	postTitle = Post.NewPostTitle("")
	if postTitle == nil { t.Error("PostTitle Construct Validate Failed...") }

	// 等価チェック
	postTitle = Post.NewPostTitle("test")
	var sameTitle = Post.NewPostTitle("test")
	if !postTitle.Equal(sameTitle) { t.Error("PostTitle Equal Test Failed...") }
	var diffTitle = Post.NewPostTitle("test2")
	if postTitle.Equal(diffTitle) { t.Error("PostTitle Not Equal Test Failed...") }

	t.Log("PostTitle Test Success")
}
