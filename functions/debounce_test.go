package functions

import (
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {

	i := 0
	h := func() {
		increment(&i)
	}
	df := Debounce(h, 1000)
	df.Call()
	time.Sleep(3 * time.Second)
	df.Cancel()
	if i != 3 {
		t.Fail()
	}
}
func increment(i *int) {
	*i += 1
}
