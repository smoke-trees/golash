package function

import (
	"testing"
)

func TestAfter(t *testing.T) {
	i := 0
	h := func() {
		increment(&i)
	}
	df := After(h, 3)

	for j := 0; j < 5; j++ {
		df.Call()
	}
	df.Reset()
	for j := 0; j < 5; j++ {
		df.Call()
	}
	if i != 6 {
		t.Fail()
	}

	df1 := After(h, 0)

	df1.Call()

	if i != 7 {
		t.Fail()
	}
}
