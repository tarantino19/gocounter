package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := 0

	for _, v := range data {
		if v == ' ' {
			fmt.Println("space detected")
			wordCount++
		}
	}

	wordCount++

	fmt.Println("word count:", wordCount)

}
