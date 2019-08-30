package string

import "strings"

// Converts the first character of s to upper case and the remaining to lower case.
func Capitalize(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
