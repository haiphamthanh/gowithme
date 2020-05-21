package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var postById map[int]*Post
var postsByAuthor map[string][]*Post

func store(post Post) {
	postById[post.Id] = &post
	postsByAuthor[post.Author] = append(postsByAuthor[post.Author], &post)
}

func main() {
	postById = make(map[int]*Post)
	postsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello world!", Author: "Sau sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Gretting Earthlings!", Author: "Shau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(postById[1])
	fmt.Println(postById[2])

	for _, post := range postsByAuthor["Shau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range postsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
