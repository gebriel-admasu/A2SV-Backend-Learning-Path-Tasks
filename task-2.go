package main

import (
	"regexp"
	"strings"
)

func WordFrequencyCount(input string) map[string]int {
	re := regexp.MustCompile(`[^\w\s]`)
	cleaned := re.ReplaceAllString(input, "")
	cleaned = strings.ToLower(cleaned)
	words := strings.Fields(cleaned)

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

func IsPalindrome(input string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := re.ReplaceAllString(input, "")
	cleaned = strings.ToLower(cleaned)

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}
