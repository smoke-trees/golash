package string

import (
	"testing"
)

var comboMarks = []string{"̀", "́", "̂", "̃", "̄", "̅", "̆", "̇", "̈", "̉", "̊", "̋", "̌", "̍", "̎", "̏", "̐", "̑", "̒", "̓", "̔", "̕", "̖", "̗", "̘", "̙", "̚", "̛", "̜", "̝", "̞", "̟", "̠", "̡", "̢", "̣", "̤", "̥", "̦", "̧", "̨", "̩", "̪", "̫", "̬", "̭", "̮", "̯", "̰", "̱", "̲", "̳", "̴", "̵", "̶", "̷", "̸", "̹", "̺", "̻", "̼", "̽", "̾", "̿", "̀", "́", "͂", "̓", "̈́", "ͅ", "͆", "͇", "͈", "͉", "͊", "͋", "͌", "͍", "͎", "͏", "͐", "͑", "͒", "͓", "͔", "͕", "͖", "͗", "͘", "͙", "͚", "͛", "͜", "͝", "͞", "͟", "͠", "͡", "͢", "ͣ", "ͤ", "ͥ", "ͦ", "ͧ", "ͨ", "ͩ", "ͪ", "ͫ", "ͬ", "ͭ", "ͮ", "ͯ", "︠", "︡", "︢", "︣",
}

func TestDeburr(t *testing.T) {
	for k, v := range deburredLetters {
		if ans := Deburr(k); ans != v {
			t.Fatalf("Expected %v, got %v", v, ans)
		}
	}

	if Deburr("×") != "×" || Deburr("÷") != "÷" {
		t.Fatalf("Deburr must ignore Latin mathematical symbols")
	}

	for _, c := range comboMarks {
		if ans := Deburr("e" + c + "i"); ans != "ei" {
			t.Fatalf("Expected ei, got %v", ans)
		}
	}
}
