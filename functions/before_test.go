package functions

import (
	"testing"
)

func TestBefore(t *testing.T) {
	i := 0
	h := func() {
		increment(&i)
	}
	df := Before(h, 2)

	for j := 0; j < 6; j++ {
		df.Call()
	}
	t.Log(i)
	df.Reset()
	for j := 0; j < 6; j++ {
		df.Call()
	}
	t.Log(i)
	if i != 4 {
		t.Fail()
	}
}
