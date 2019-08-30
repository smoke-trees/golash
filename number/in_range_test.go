package number

import "testing"

func TestInRange(t *testing.T) {
	if !InRange(3, 2, 4) {
		t.Fail()
	}

	if !InRange(-3, -2, -6) {
		defer func () {
			recover()
		}()
		t.Fail()
	}
}

func TestInRange2(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("The code did not panic")
		}
	}()

	InRange(5.2, 0.0, 4)
}
