package Post

// 投稿データ
type Post struct {
	title string
}

// コンストラクタ
func NewPost(title string) *Post {
	var post *Post = new(Post)
	post.title = title
	return post
}

// タイトル取得
func (this *Post) Title() string {
	return this.title
}
