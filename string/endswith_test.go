package string

import (
	"math"
	"testing"
)

func TestEndsWith(t *testing.T) {
	if !EndsWith("abc", "c", -1) {
		t.Fatalf("should treat negative position as 0")
	}

	if !EndsWith("abc", "c", 0) {
		t.Fatalf("should return true if a string ends with the target")
	}

	if EndsWith("abc", "b", 0) {
		t.Fatalf("should return false if a string does not end with the target")
	}

	if !EndsWith("abc", "b", 2) {
		t.Fatalf("Should work with a position")
	}

	if !EndsWith("abc", "c", math.MaxInt64) {
		t.Fatalf("Should work with a position >= length")
	}

}
