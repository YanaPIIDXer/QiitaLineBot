package post

// 投稿データ
type Post struct {
	Title string
}

// コンストラクタ
func NewPost(title string) *Post {
	var post *Post = new(Post)
	post.Title = title
	return post
}
