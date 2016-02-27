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

var step1a_rules = []PorterRule{
	make_literal_substitution("SSES", "SS"),
	make_literal_substitution("IES", "I"),
	make_literal_substitution("SS", "SS"),
	make_literal_substitution("S", ""),
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
