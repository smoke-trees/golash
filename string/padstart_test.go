package string

import "testing"

func TestPadStart(t *testing.T) {
	if ans := PadStart("abc", 6, " "); ans != "   abc" {
		t.Fatalf("It should pad a string to a given length, got %v", ans)
	}

	if ans := PadStart("abc", 6, "_-"); ans != "_-_abc" {
		t.Fatalf("It should truncate pad characters to fit the pad length, got %v", ans)
	}
}

func TestPadStart2(t *testing.T) {
	if ans := PadStart("abc", -1, " "); ans != "abc" {
		t.Fatalf("It should pad a string to a given length, got %v", ans)
	}
}
