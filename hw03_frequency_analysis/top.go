package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

var valuesToReturn = 10

func splitText(text string) []string {
	replaceCharacters := []string{",", ".", ",", "?", "!"}

	text = strings.ReplaceAll(text, "\n", " ")
	for _, char := range replaceCharacters {
		text = strings.ReplaceAll(text, char, "")
	}

	words := strings.Split(strings.ToLower(text), " ")
	return words
}

func sortWordsCounter(counter map[string]int) []string {
	result := []string{}
	type wordStats struct {
		word  string
		count int
	}
	words := []wordStats{}

	for k, v := range counter {
		words = append(words, wordStats{k, v})
	}

	sort.Slice(words, func(i, j int) bool { return words[i].count > words[j].count })
	for index, v := range words {
		if index >= valuesToReturn {
			break
		}
		if v.word != "" { // edge case. If input text is empty
			result = append(result, v.word)
		}
	}
	return result
}

// Top10 function returns 10 most used words in text
func Top10(text string) []string {
	words := splitText(text)
	wordsCounter := map[string]int{}
	for _, word := range words {
		wordsCounter[word]++
	}

	return sortWordsCounter(wordsCounter)
}
