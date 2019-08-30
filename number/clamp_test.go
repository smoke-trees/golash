package number

import "testing"

func TestClamp(t *testing.T) {
	// uint test
	if Clamp(5, 7, 10) != 7 {
		t.Fail()
	}

	// int test
	if Clamp(-10, -5, 5) != -5 {
		t.Fail()
	}

	// float test
	if Clamp(-15.5, -2.4, 5.7) != -2.4 {
		t.Fail()
	}
}
