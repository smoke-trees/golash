package string

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	if Repeat("*", 3) != "***" {
		t.Fatalf("It should repeat a string `n` times")
	}

	if Repeat("abc", 0) != "" {
		t.Fatalf("It should return an empty string if `n` is <= 0")
	}
}
