package gotext

import (
	"testing"
)

func is_vowel(t *testing.T, r rune) {
	if is_standard_consonant(r) {
		t.Fatalf("rune %s was thought to be a consonant", string(r))
	}
}

func TestRuneContainsVowels(t *testing.T) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for i := range vowels {
		is_vowel(t, vowels[i])
	}
}

func is_consonant(t *testing.T, r rune) {
	if !is_standard_consonant(r) {
		t.Fatalf("rune %s was thought to be a vowel", string(r))
	}
}

func TestRuneContainsConsonant(t *testing.T) {
	is_consonant(t, 'c')
	is_consonant(t, 'C')

	is_consonant(t, 'Z')
	is_consonant(t, 'Z')

	is_consonant(t, 'p')
	is_consonant(t, 'P')
}

func TestEmptyStringIsVowelLess(t *testing.T) {
	if contains_vowel("") {
		t.Fatalf("empty string was thought to be a vowel..")
	}
}

func TestVowelContainingString(t *testing.T) {
	if !contains_vowel("a") {
		t.Fatalf("'a' was said to not contain a vowel")
	}

	if !contains_vowel("bat") {
		t.Fatalf("'bat' was said to not contain a vowel")
	}

	if !contains_vowel("apple") {
		t.Fatalf("'apple' was said to not contain a vowel")
	}

	if !contains_vowel("ba") {
		t.Fatalf("'ba' was said to not contain a vowel")
	}

	if !contains_vowel("cy") {
		t.Fatalf("'cy' has a vowel (y)")
	}
}

func TestYVowelness(t *testing.T) {
	if contains_vowel("y") {
		t.Fatalf("'y' is not a vowel on its own")
	}

	if !contains_vowel("my") {
		t.Fatalf("'my' contains a vowel")
	}

	if !contains_vowel("ay") {
		t.Fatalf("'ay' contains a vowel (a)")
	}

	if !contains_vowel("ya") {
		t.Fatalf("'YA' contains a vowel (A)")
	}
}

func TestStringWithoutVowels(t *testing.T) {
	if contains_vowel("z") {
		t.Fatalf("'z' was said to contain a vowel")
	}

	if contains_vowel("zbc") {
		t.Fatalf("'zbc' was said to contain a vowel")
	}

	if contains_vowel("zzz") {
		t.Fatalf("'z' was said to contain a vowel")
	}
}
