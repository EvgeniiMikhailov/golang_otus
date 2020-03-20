package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

type wordStats struct {
	word  string
	count int
}

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

func wordsCounterToSortedSlice(counter map[string]int) []wordStats {
	words := []wordStats{}

	for k, v := range counter {
		words = append(words, wordStats{k, v})
	}

	sort.Slice(words, func(i, j int) bool { return words[i].count > words[j].count })
	return words
}

// Top10 function returns 10 most used words in text
func Top10(text string) []string {
	words := splitText(text)
	wordsCounter := map[string]int{}
	for _, word := range words {
		wordsCounter[word]++
	}

	wordsWithFrequency := wordsCounterToSortedSlice(wordsCounter)
	result := []string{}
	for index, v := range wordsWithFrequency {
		if index >= valuesToReturn {
			break
		}
		if v.word != "" { // edge case. If input text is empty
			result = append(result, v.word)
		}
	}
	return result
}
