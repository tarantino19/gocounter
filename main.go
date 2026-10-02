package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "./bigtext.txt"
	log.SetFlags(0)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("failed to read file: ", err)
	}

	wordCount := CountWords(file)
	fmt.Println(wordCount)
}

func CountWords(file io.Reader) int {
	//io.Reader for string file bytes + other data streams support
	wordCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to scan file: ", err)
	}

	return wordCount
}
