package gotext

import (
	"strings"
	"testing"
)

func match_against_rules(t *testing.T, str string, rules []PorterRule, ans string) {
	s := strings.ToLower(str)
	a := strings.ToLower(ans)
	r := strings.ToLower(PorterStemRules(s, rules))

	if r != a {
		t.Fatalf("str=%s != ans=%s (res=%s)\n", s, a, r)
	}
}

func TestPorterEmptyString(t *testing.T) {

}

func step1a_check(t *testing.T, str string, ans string) {
	match_against_rules(t, str, step1a_rules, ans)
}

func TestPorterStep1a(t *testing.T) {
	step1a_check(t, "CARESSES", "CARESS")
	step1a_check(t, "ponies", "poni")
	step1a_check(t, "ties", "ti")
	step1a_check(t, "CARESS", "CARESS")
	step1a_check(t, "cats", "cat")
}

func step1b_check(t *testing.T, str string, ans string) {
	match_against_rules(t, str, step1b_rules, ans)
}

func TestPorterStep1b(t *testing.T) {
	step1b_check(t, "FEED", "FEED")
	step1b_check(t, "feed", "feed")

	step1b_check(t, "AGREED", "AGREE")
	step1b_check(t, "agreed", "agree")
}
