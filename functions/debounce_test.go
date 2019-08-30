package function

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
		t.Fatalf("Error in Debounce %d", i)
	}
}
func increment(i *int) {
	*i += 1
}
