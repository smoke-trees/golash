package string

import "testing"

func TestEscapeRegExp(t *testing.T) {
	escaped := "\\^\\$\\.\\*\\+\\?\\(\\)\\[\\]\\{\\}\\|\\\\"
	unescaped := "^$.*+?()[]{}|\\"

	if EscapeRegExp(unescaped+unescaped) != escaped+escaped {
		t.Fatalf("It should escape values")
	}

	if EscapeRegExp("abc") != "abc" {
		t.Fatalf("It should handle strings with nothing to escape")
	}
}
