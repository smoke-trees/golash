package string

import (
	"regexp"
)

func EscapeRegExp(s string) string {
	return regexp.MustCompile(`[\\^$.*+?()[\]{}|]`).ReplaceAllStringFunc(s, func(m string) string {
		return "\\" + m
	})
}
