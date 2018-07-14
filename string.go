package ut

import (
	"strings"
	"unicode"
)

// Reverse reverse a string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// TrimSpace remove all spaces in the string.
func TrimSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// LimitLength limit the maximum length of the string.
func LimitLength(s string, length int, ellipses string) string {
	str := []rune(s)
	if len(str) > length {
		return string(str[:length-len(ellipses)]) + ellipses
	}
	return s
}
