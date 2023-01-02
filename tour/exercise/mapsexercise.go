package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount is exported
func WordCount(s string) map[string]int {
	stringSplit := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, value := range stringSplit {

		if wordCount[value] == 0 {
			wordCount[value] = 1
		} else if wordCount[value] >= 1 {
			wordCount[value] = wordCount[value] + 1
		}
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
