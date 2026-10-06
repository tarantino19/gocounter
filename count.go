package main

import (
	"bufio"
	"io"
	"os"
)

type Counts struct {
	Lines int
	Words int
	Bytes int
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)

	if err != nil {
		return Counts{}, err
	}

	defer file.Close()

	const offsetStart = 0

	byteCount := CountBytes(file)
	file.Seek(offsetStart, io.SeekStart)
	wordCount := CountWords(file)
	file.Seek(offsetStart, io.SeekStart)
	lineCount := CountLines(file)

	return Counts{
		Bytes: byteCount,
		Words: wordCount,
		Lines: lineCount,
	}, nil

}

func CountWords(file io.Reader) int {
	//io.Reader for string file bytes + other data streams support
	wordCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if scanner.Err() != nil {
		return 0
	}

	return wordCount
}

func CountLines(r io.Reader) int {
	linesCount := 0

	reader := bufio.NewReader(r)

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			break
		}

		if r == '\n' {
			linesCount++
		}
	}

	return linesCount
}

func CountBytes(r io.Reader) int {
	byteCount, _ := io.Copy(io.Discard, r)

	return int(byteCount)

}
