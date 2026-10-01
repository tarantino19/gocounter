package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./wordsd.txt")

	log.SetFlags(0)
	if err != nil {
		log.Fatalln("failed to read file: ", err)
	}

	wordCount := CountWords(data)
	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := len(bytes.Fields(data))
	return words
}
