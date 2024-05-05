package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {

	s = strings.ToLower(s)
	alphabet := " abcdefghijklmnopqrstuvwxyz"

	for _, letter := range alphabet {
		if !strings.Contains(s, string(letter)) {
			return "not pangram"
		}
	}

	return "pangram"
}

func main() {

	words := []string{
		"The quick brown fox jumps over the lazy dog",
		"We promptly judged antique ivory buckles for the next prize",
		"We promptly judged antique ivory buckles for the prize",
	}

	for _, word := range words {
		fmt.Println(pangrams(word))
	}
}
