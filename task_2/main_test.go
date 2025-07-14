package main

import (
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	tests := []struct {
		text     string
		expected map[string]int
	}{
		{
			"Hello, world! Hello from go.",
			map[string]int{
				"hello": 2,
				"world": 1,
				"go":    1,
				"from":  1,
			},
		},
		{
			"Go is great.",
			map[string]int{
				"go":    1,
				"is":    1,
				"great": 1,
			},
		},
		{
			"Hello! hello? HELLO!!",
			map[string]int{
				"hello": 3,
			},
		},
		{
			"Hello",
			map[string]int{
				"hello": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			result := WordFrequencyCount(tt.text)
			for word, count := range tt.expected {
				if result[word] != count {
					t.Errorf("expected %d for word '%s', got %d", count, word, result[word])
				}
			}
		})
	}
}

func TestPalindromeCheck(t *testing.T) {
	tests := []struct {
		text     string
		expected bool
	}{
		{"A man, a plan, a canal, Panama!", true},
		{"No 'x' in Nixon", true},
		{"Hello, world!", false},
		{"Racecar", true},
		{"Was it a car or a cat I saw?", true},
		{"Not a palindrome", false},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			result := PalindromeCheck(tt.text)
			if result != tt.expected {
				t.Errorf("expected %v for text '%s', got %v", tt.expected, tt.text, result)
			}
		})
	}
}
