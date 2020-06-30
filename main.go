package main

import (
	"fmt"

	"github.com/YanaPIIDXer/QiitaLineBot/src/domain/post"
)

func main() {
	var post *post.Post = post.NewPost("MyPost")
	fmt.Println(post.Title)
}
