package main

import (
	"bytes"
	"fmt"
	"regexp"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) (mostCommon string) {

	var currentWord string
	var mostCommonCount int
	banedWords := make(map[string]bool, len(banned))
	wordCount := make(map[string]int)
	wordEndings := regexp.MustCompile(`(\.|'|,|;|!|\?|\s)+`)

	for _, ban := range banned {
		lowered := bytes.ToLower([]byte(ban))
		banedWords[string(lowered)] = true
	}

	for i := 0; i < len(paragraph); i++ {
		bInStr := fmt.Sprintf("%c", unicode.ToLower(rune(paragraph[i])))
		if wordEndings.MatchString(bInStr) {
			if len(currentWord) == 0 {
				continue
			}

			if _, isBanned := banedWords[currentWord]; isBanned {
				currentWord = ""
				continue
			}

			count, _ := wordCount[currentWord]
			wordCount[currentWord] = count + 1
			if mostCommonCount < count+1 {
				mostCommon = currentWord
				mostCommonCount = count + 1
			}
			currentWord = ""
		} else {
			currentWord = fmt.Sprintf("%s%s", currentWord, bInStr)
			if i == len(paragraph)-1 {
				count, _ := wordCount[currentWord]
				wordCount[currentWord] = count + 1
				if mostCommonCount < count+1 {
					mostCommon = currentWord
					mostCommonCount = count + 1
				}
			}
		}

	}

	return
}

func main() {
	fmt.Printf("most common word is as follows '%s'\n", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"Hit"}))
}
