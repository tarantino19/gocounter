package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./words.txt")
	fmt.Println("Hello, Words!")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	wordCount := CountWords(data)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := len(bytes.Fields(data))
	return words
}
