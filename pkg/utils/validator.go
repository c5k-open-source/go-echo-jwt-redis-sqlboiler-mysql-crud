package utils

import (
	"fmt"
	"regexp"
	"unicode"
)

func ValidateEmailInput(em string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(em)
}

func ValidatePasswordInput(p string) error {
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range p {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("Password did not match with rules.\n Password must have at least one %s character.", name)
	}
	return nil
}

func ValidateDBName(n string) bool {
	re := regexp.MustCompile("^[a-z0-9_-]*$")
	return re.MatchString(n)
}
