package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"fmt"
	"strings"
)

func splitText(text string) []string {
	replaceCharacters := []string{",", ".", ",", "?", "!"}

	text = strings.ReplaceAll(text, "\n", " ")
	for _, char := range replaceCharacters {
		text = strings.ReplaceAll(text, char, "")
	}

	words := strings.Split(strings.ToLower(text), " ")
	return words
}

func Top10(text string) []string {
	words := splitText(text)
	fmt.Println(words)
	return nil
}
