package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello world",
		Author: Author{
			Id:   "1",
			Name: "Shau sheng new",
		},
	}

	// Create file
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}

	// Encode
	encode := xml.NewEncoder(xmlFile)
	encode.Indent("", "\t")
	err = encode.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}
}
