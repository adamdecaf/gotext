package gotext

import (
	"fmt"
)

var vowels = []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}

func is_standard_consonant(r rune) bool {
	for i := range vowels {
		if r == vowels[i] {
			return false
		}
	}
	return true
}

func contains_vowel(s string) bool {
	if len(s) == 1 {
		rs := []rune(s)
		r := rs[0]
		return !is_standard_consonant(r)
	}

	runes := []rune(s)
	for i := range runes {
		if i == 0 {
			continue
		}
		r1 := runes[i-1]
		r2 := runes[i]

		fmt.Printf("r1=%s, r2=%s\n", string(r1), string(r2))

		if !is_standard_consonant(r1) || !is_standard_consonant(r2) {
			fmt.Printf("ooo -- r1=%s, r2=%s\n", string(r1), string(r2))
			return true
		}

		if r2 == 'y' && is_standard_consonant(r1) {
			fmt.Println("OOO -- r1=%s, r2=%s", string(r1), string(r2))
			return true
		}
	}

	return false
}
