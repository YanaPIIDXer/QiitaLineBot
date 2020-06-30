package main

import (
	"fmt"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/Post"
)

func main() {
	var post *Post.Post = Post.NewPost("MyPost")
	fmt.Println(post.Title())
}
