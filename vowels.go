package main

import (
	"golang.org/x/exp/slices"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func CountVowels(s string) int {
	var vowelsCount int
	for i := range s {
		if slices.Contains(vowels, s[i]) {
			vowelsCount++
		}
	}

	return vowelsCount
}
