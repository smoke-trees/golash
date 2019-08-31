package string

import (
	"math"
	"testing"
)

func TestStartsWith(t *testing.T) {
	if !StartsWith("abc", "a", 0) {
		t.Fatalf("It should return `true` if a string starts with `target")
	}

	if StartsWith("abc", "b", 0) {
		t.Fatalf("should return `false` if a string does not start with `target")
	}
}

func TestStartsWith2(t *testing.T) {
	if StartsWith("abc", "a", math.MaxInt64) {
		t.Fatalf("It should work with `position` >= `length`")
	}

	if !StartsWith("abc", "a", -1) {
		t.Fatalf("It should work with `position` >= `length`")
	}
}
