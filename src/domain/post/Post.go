package Post

// 投稿データ
type Post struct {
	title *PostTitle		// タイトル
}

// コンストラクタ
func NewPost(title *PostTitle) *Post {
	if title == nil { return nil }
	
	var post *Post = new(Post)
	post.title = title
	
	return post
}
