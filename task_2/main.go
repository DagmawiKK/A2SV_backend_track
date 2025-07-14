package main

import (
	"regexp"
	"strings"
)

func WordFrequencyCount(text string) map[string]int {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func PalindromeCheck(text string) bool {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\w]`)
	text = re.ReplaceAllString(text, "")

	left, right := 0, len(text)-1

	for left < right {
		if text[left] != text[right] {
			return false
		}
		left++
		right--
	}

	return true
}
