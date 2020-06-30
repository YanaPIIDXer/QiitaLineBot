package main

import (
	"fmt"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/post"
)

func main() {
	var post *Post.Post = Post.NewPost(Post.NewPostTitle("MyPost"))
	fmt.Println(post.Title().Value())
}
