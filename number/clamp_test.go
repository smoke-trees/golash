package number

import "testing"

func TestClamp(t *testing.T) {
	// uint test
	if Clamp(uint(5), uint(7), uint(10)) != uint(7) {
		t.Fail()
	}

	// uint8 test
	if Clamp(uint8(5), uint8(7), uint8(10)) != uint8(7) {
		t.Fail()
	}

	// uint16 test
	if Clamp(uint16(5), uint16(7), uint16(10)) != uint16(7) {
		t.Fail()
	}

	// uint32 test
	if Clamp(uint32(5), uint32(7), uint32(10)) != uint32(7) {
		t.Fail()
	}

	// uint64 test
	if Clamp(uint64(5), uint64(7), uint64(10)) != uint64(7) {
		t.Fail()
	}

	// int test
	if Clamp(-10, -5, 5) != -5 {
		t.Fail()
	}

	// int8 test
	if Clamp(int8(5), int8(7), int8(10)) != int8(7) {
		t.Fail()
	}

	// int16 test
	if Clamp(int16(5), int16(7), int16(10)) != int16(7) {
		t.Fail()
	}

	// int32 test
	if Clamp(int32(5), int32(7), int32(10)) != int32(7) {
		t.Fail()
	}

	// int64 test
	if Clamp(int64(5), int64(7), int64(10)) != int64(7) {
		t.Fail()
	}

	// float32 test
	if Clamp(float32(-15.5), float32(-2.4), float32(5.7)) != float32(-2.4) {
		t.Fail()
	}

	// float64 test
	if Clamp(-15.5, -2.4, 5.7) != -2.4 {
		t.Fail()
	}
}

func TestClamp2(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	Clamp("n", "l", "u")
}

func TestClamp3(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()
	Clamp(4, 5.1, 7)
}
