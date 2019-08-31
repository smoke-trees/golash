package string

import "testing"

func TestPadEnd(t *testing.T) {
	if ans := PadEnd("abc", 6, " "); ans != "abc   " {
		t.Fatalf("It should pad a string to a given length, got %v", ans)
	}

	if ans := PadEnd("abc", 6, "_-"); ans != "abc_-_" {
		t.Fatalf("It should truncate pad characters to fit the pad length, got %v", ans)
	}
}

func TestPadEnd2(t *testing.T) {
	if ans := PadEnd("abc", -1, " "); ans != "abc" {
		t.Fatalf("It should pad a string to a given length, got %v", ans)
	}
}
