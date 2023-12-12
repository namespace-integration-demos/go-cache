package main

import "testing"

func TestCountVowels(t *testing.T) {

	inputs := map[string]int{
		"hello": 2,
		"world": 1,
	}

	for s, expected := range inputs {
		if c := CountVowels(s); c != expected {
			t.Errorf("Expected %d vowels, counted %d.", expected, c)
		}
	}
}
