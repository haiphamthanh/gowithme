package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id      int       `json:"id"`
	Content string    `json:"content"`
	Author  Author    `json:"author"`
	Comment []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer file.Close()

	// Create decoder
	decoder := json.NewDecoder(file)

	// Use decoder to write data to post item
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding Json: ", err)
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return
	}
	defer file.Close()

	// Read file
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
		return
	}

	// Unmarshal file to variable
	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	_, err := decode("post.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
