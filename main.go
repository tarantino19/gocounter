package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("error: no filename specified")
	}

	total := 0
	fileNames := os.Args[1:]

	for _, filename := range os.Args[1:] {
		wordCount := CountWordsInFile(filename)
		total += wordCount
		fmt.Println(filename, wordCount)
	}

	if len(fileNames) > 1 {
		fmt.Println("total word count:", total)
	}
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln("failed to read file: ", err)
	}

	return CountWords(file)

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
