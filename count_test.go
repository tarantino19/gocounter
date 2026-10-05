package main_test

import (
	"strings"
	"testing"

	counter "github.com/tarantino19/gocounter"
)

func TestCountWords(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "This is a sentence.  This is another ",
			wants: 7,
		},
		{
			name:  "prefixed multiple spaces",
			input: "    Hello",
			wants: 1,
		},
		{
			name:  "suffixed multiple spaces",
			input: "Hello  ",
			wants: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := counter.CountWords(r)

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	//lines based on wc man pages func
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple five words, 1 new line",
			input: "one two three for five\n",
			wants: 1,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "no new lines",
			input: "one two three four five",
			wants: 0,
		},
		{
			name:  "no new line at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name:  "multi new line strings",
			input: "\n\n\n\n",
			wants: 4,
		},
		{
			name:  "multi 5 line strings",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := counter.CountLines(r)

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}
		})
	}

}
