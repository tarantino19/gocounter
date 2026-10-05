package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	fileNames := os.Args[1:]
	didError := false

	for _, filename := range fileNames {
		wordCount, err := CountWordsInFile(filename)

		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "FAIL:", err)
			continue
		}

		total += wordCount
		fmt.Println(wordCount, ":", filename)
	}

	if len(fileNames) == 0 {
		wordCount := CountWords(os.Stdin)
		fmt.Println(wordCount, "hello")
	}

	if len(fileNames) > 1 {
		fmt.Println(total, ":total word count")
	}

	if didError {
		os.Exit(1)
	}
}
