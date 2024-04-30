package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Book struct {
	Title         string `json:"title"`
	Artist        string `json:"artist"`
	Age           int    `json:"age,omitempty"`
	TotalPages    int    `json:"total_pages"`
	FavoritePages []int  `json:"favorite_pages"`
	Color         string `json:"color"`
	IsActive      bool   `json:"is_active"`
}

func main() {

	// Read json file:
	fileBytes, err := os.ReadFile("file.json")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(fileBytes))

	// Parse data from json file into an array of type Book
	var books []Book

	if err := json.Unmarshal(fileBytes, &books); err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
