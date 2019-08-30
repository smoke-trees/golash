package number

import "testing"

func TestInRange(t *testing.T) {
	// test for int
	if !InRange(3, 2, 4) {
		t.Fail()
	}

	// test for int8
	if !InRange(int8(3), int8(2), int8(4)) {
		t.Fail()
	}

	// test for int16
	if !InRange(int16(3), int16(2), int16(4)) {
		t.Fail()
	}

	// test for int32
	if !InRange(int32(3), int32(2), int32(4)) {
		t.Fail()
	}

	// test for int64
	if !InRange(int64(3), int64(2), int64(4)) {
		t.Fail()
	}

	// test for uint
	if !InRange(uint(3), uint(2), uint(4)) {
		t.Fail()
	}

	// test for uint8
	if !InRange(uint8(3), uint8(2), uint8(4)) {
		t.Fail()
	}

	// test for uint16
	if !InRange(uint16(3), uint16(2), uint16(4)) {
		t.Fail()
	}

	// test for uint32
	if !InRange(uint32(3), uint32(2), uint32(4)) {
		t.Fail()
	}

	// test for uint64
	if !InRange(uint64(3), uint64(2), uint64(4)) {
		t.Fail()
	}

	// test for float32
	if !InRange(float32(3), float32(2), float32(4)) {
		t.Fail()
	}

	if !InRange(3.4, 2.8, 4.2) {
		t.Fail()
	}
}

// Test for homogeneity of arguments
func TestInRange2(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()

	InRange(5.2, 0.0, 4)
}

// Test for non-numeric values
func TestInRange3(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()

	InRange("s", 4.2, 6.9)
}
