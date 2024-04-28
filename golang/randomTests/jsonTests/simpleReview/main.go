package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type book struct {
	Title string `json:"title"`
}

func main() {

	// var jsonByte = []byte(`[{"title": "Caderno do Felipe"},{"title": "Caderno da Daiana"}]`)
	// fmt.Println(string(jsonByte))

	var jsonString = `
	[
		{"title": "Caderno do Felipe"},
		{"title": "Caderno da Daiana"}
	]`
	fmt.Println(jsonString)

	var books []book

	if err := json.Unmarshal([]byte(jsonString), &books); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", books)
}
