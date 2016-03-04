package gotext

import (
	"fmt"
	"regexp"
	"strings"
)

type PorterRule interface {
	Apply(str string) (string, error)
}

type LiteralSubstitution struct {
	PorterRule

	Find string
	Replace string
}

func (p LiteralSubstitution) Apply(str string) (string, error) {
	r := regexp.MustCompile("(?im)" + p.Find)
	res := r.ReplaceAllLiteralString(strings.ToLower(str), p.Replace)
	if (len(res) == len(str)) && (len(p.Find) != len(p.Replace)) {
		return "", fmt.Errorf("str=%s, res=%s", str, res)
	}
	return res, nil
}

func make_literal_substitution(find string, replace string) LiteralSubstitution {
	sub := LiteralSubstitution{}
	sub.Find = strings.ToLower(find)
	sub.Replace = strings.ToLower(replace)
	return sub
}

type NonNullWordLiteralSubstitution struct {
	PorterRule

	Suffix string
	Replace string
}

func (nn NonNullWordLiteralSubstitution) Apply(str string) (string, error) {
	s := strings.ToLower(str)
	suf := strings.ToLower(nn.Suffix)

	if strings.HasSuffix(s, suf) {
		parts := strings.Split(s, suf)
		if len(parts) > 1 {
			prefix := strings.Join(parts, "")
			// only 'non-null' words qualify for replacement
			if contains_vowel(prefix) {
				return prefix + nn.Replace, nil
			}
			return str, nil
		}

	}
	return str, nil
}

func make_non_null_literal_sub(suffix, replace string) NonNullWordLiteralSubstitution {
	nn := NonNullWordLiteralSubstitution{}
	nn.Suffix = strings.ToLower(suffix)
	nn.Replace = strings.ToLower(replace)
	return nn
}

var step1a_rules = []PorterRule{
	make_literal_substitution("SSES", "SS"),
	make_literal_substitution("IES", "I"),
	make_literal_substitution("SS", "SS"),
	make_literal_substitution("S", ""),
}

var step1b_rules = []PorterRule{
	make_non_null_literal_sub("EED", "EE"),
}

var default_porter_stem_rules =
	append([]PorterRule{}, step1a_rules...)

func PorterStem(str string) string {
	return PorterStemRules(str, default_porter_stem_rules)
}

func PorterStemRules(str string, rules []PorterRule) string {
	var ans string
	ans = str
	for i := range rules {
		r, err := rules[i].Apply(ans)
		if err == nil {
			ans = r
			if r != "" {
				return ans
			}
		}
	}
	return ans
}
