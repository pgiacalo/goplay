package main

// KEY POINT: structs using reflection for marshalling/unmarshalling must have exported fields (i.e., Capitalized names)

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
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

func main() {

	post := Post{Id: 1, Content: "Hello World!"}
	post.Author = Author{Id: 2, Name: "Sau Sheong"}
	post.Comments = make([]Comment, 2)	//make an array of Comment structs of length 2
	post.Comments[0] = Comment{Id: 1, Content: "Have a great day!", Author: "Adam"}
	post.Comments[1] = Comment{Id: 2, Content: "How are you today?", Author: "Betty"}

	// --- The Post object can also be created as follows (but this seems a lot less intuitive) ---
	// --- Either way, the resulting json file produced is identical ---

	//post := Post{
	//	Id:      1,
	//	Content: "Hello World!",
	//	Author: Author{
	//		Id:   2,
	//		Name: "Sau Sheong",
	//	},
	//	Comments: []Comment{
	//		Comment{
	//			Id:      1,
	//			Content: "Have a great day!",
	//			Author:  "Adam",
	//		},
	//		Comment{
	//			Id:      2,
	//			Content: "How are you today?",
	//			Author:  "Betty",
	//		},
	//	},
	//}

	jsonFile, err := os.Create("post1.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}

// Produces post.json file containing the following:

// {"id":1,"content":"Hello World!","author":{"id":2,"name":"Sau Sheong"},"comments":[{"id":1,"content":"Have a great day!","author":"Adam"},{"id":2,"content":"How are you today?","author":"Betty"}]}
