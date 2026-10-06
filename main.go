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
		counts, err := CountFile(filename)

		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "FAIL:", err)
			continue
		}

		total += counts.Words
		fmt.Println(counts.Bytes, counts.Words, counts.Lines, ":", filename)
	}

	if len(fileNames) == 0 {
		wordCount := CountWords(os.Stdin)
		byteCount := CountBytes(os.Stdin)
		lineCount := CountLines(os.Stdin)
		fmt.Println(wordCount, byteCount, lineCount)
	}

	if len(fileNames) > 1 {
		fmt.Println(total, ":total word count")
	}

	if didError {
		os.Exit(1)
	}
}
