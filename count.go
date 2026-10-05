package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err) //error wrappin
	}

	defer file.Close()

	return CountWords(file), nil

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
		return 0
	}

	return wordCount
}
