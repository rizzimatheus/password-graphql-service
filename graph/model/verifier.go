package model

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

// VarifyRules returns which 'rules' have not been satisfied by the 'password'
func VarifyRules(password string, rules []*Rule) ([]*string, error) {
	nomatch := []*string{}
	for _, rule := range rules {
		switch rule.Rule {
		case "minSize":
			if !MinSize(password, rule.Value) {
				nomatch = append(nomatch, &rule.Rule)
			}
		case "minUppercase":
			if !MinUppercase(password, rule.Value) {
				nomatch = append(nomatch, &rule.Rule)
			}
		case "minLowercase":
			if !MinLowercase(password, rule.Value) {
				nomatch = append(nomatch, &rule.Rule)
			}
		case "minDigit":
			if !MinDigit(password, rule.Value) {
				nomatch = append(nomatch, &rule.Rule)
			}
		case "minSpecialChars":
			if !MinSpecialChars(password, rule.Value) {
				nomatch = append(nomatch, &rule.Rule)
			}
		case "noRepeted":
			if !NoRepeted(password) {
				nomatch = append(nomatch, &rule.Rule)
			}
		default:
			errMsg := fmt.Sprintf("Error: '%s' is an invalid rule", rule.Rule)
			return nil, errors.New(errMsg)
		}
	}

	return nomatch, nil
}

// MinSize checks if the 'password' has at least 'value' characters
func MinSize(password string, value int) bool {
	return len(password) >= value
}

// MinUppercase checks if the 'password' has at least 'value' uppercase characters
func MinUppercase(password string, value int) bool {
	count := 0
	for _, r := range password {
		if unicode.IsUpper(r) {
			count++
			if count >= value {
				return true
			}
		}
	}
	return value == 0
}

// MinLowercase checks if the 'password' has at least 'value' lowercase characters
func MinLowercase(password string, value int) bool {
	count := 0
	for _, r := range password {
		if unicode.IsLower(r) {
			count++
			if count >= value {
				return true
			}
		}
	}
	return value == 0
}

// MinDigit checks if the 'password' has at least 'value' digits
func MinDigit(password string, value int) bool {
	count := 0
	for _, r := range password {
		if unicode.IsDigit(r) {
			count++
			if count >= value {
				return true
			}
		}
	}
	return value == 0
}

// MinSpecialChars checks if the 'password' has at least 'value' special characters. Special characters: "!@#$%^&*()-+\/{}[]"
func MinSpecialChars(password string, value int) bool {
	re := regexp.MustCompile(`[!@#$%^&*()-+\/{}\[\]]`)
	count := len(re.FindAll([]byte(password), -1))
	return count >= value
}

// NoRepeted checks if the 'password' has no repeated characters in sequence. ("aab": false, "aba": true)
func NoRepeted(password string) bool {
	if len(password) <= 1 {
		return true
	}

	r := password[0]
	for i := 1; i < len(password); i++ {
		if r == password[i] {
			return false
		}

		r = password[i]
	}

	return true
}
