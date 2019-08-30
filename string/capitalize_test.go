package string

import "testing"

func TestCapitalize(t *testing.T) {
	if Capitalize("FRED") != "Fred" {
		t.Fatalf(Capitalize("FRED"))
	}

	if Capitalize("fRED") != "Fred" {
		t.Fail()
	}

	if Capitalize("fred") != "Fred" {
		t.Fail()
	}

	if Capitalize("tHiS iS a SaMpLe sEnTenCe") != "This is a sample sentence" {
		t.Fail()
	}
}
