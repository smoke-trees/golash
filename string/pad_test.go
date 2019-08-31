package string

import "testing"

func TestPad(t *testing.T) {
	if ans := Pad("abc", 6, " "); ans != " abc  " {
		t.Fatalf("It should pad a string to a given length. Got '%v'", ans)
	}

	if ans := Pad("abc", 8, " "); ans != "  abc   " {
		t.Fatalf("should truncate pad characters to fit the pad length")
	}
}

func TestPad2(t *testing.T) {
	if ans := Pad("abc", 8, "_-"); ans != "_-abc_-_" {
		t.Fatalf("should truncate pad characters to fit the pad length. Got %v", ans)
	}
}
