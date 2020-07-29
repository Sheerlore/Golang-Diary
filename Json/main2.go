package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json: "title"`
	Author string `json: "author"`
}

func main() {
	book := Book{Title: "Learning concurreny in Python", Author: "Elliot Forbes"}
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}