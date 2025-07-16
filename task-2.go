package main

import (
	"regexp"
	"strings"
)

// WordFrequencyCount returns a map with the frequency of each word in the input string.
// It ignores punctuation and is case-insensitive.
func WordFrequencyCount(input string) map[string]int {
	// Remove punctuation using regex
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

// IsPalindrome checks if the input string is a palindrome (ignoring spaces, punctuation, and case).
func IsPalindrome(input string) bool {
	// Remove non-alphanumeric characters and convert to lowercase
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
