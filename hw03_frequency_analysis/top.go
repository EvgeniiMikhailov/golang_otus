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

	splittedText := strings.Split(strings.ToLower(text), " ")
	return splittedText
}

func prettyPrint(slice []string) {
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func Top10(text string) []string {
	splittedText := splitText(text)
	fmt.Println(splittedText)
	prettyPrint(splittedText)
	return nil
}
