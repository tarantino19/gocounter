package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := countWords(data)
	fmt.Println("word count:", wordCount)
}

func countWords(data []byte) int {

	wordCount := 0

	for _, v := range data {
		if v == ' ' {
			fmt.Println("space detected")
			wordCount++
		}
	}

	wordCount++

	return wordCount
}
