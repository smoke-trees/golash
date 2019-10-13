package string

import (
	"regexp"
)

// EscapeRegExp escapes the regex in a string literal
func EscapeRegExp(s string) string {
	return regexp.MustCompile(`[\\^$.*+?()[\]{}|]`).ReplaceAllStringFunc(s, func(m string) string {
		return "\\" + m
	})
}
