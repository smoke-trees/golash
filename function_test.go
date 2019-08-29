package golash

import (
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	i := 0
	h := func() {
		i += 1
	}
	df := Debounce(h, 1000)
	df.Call()
	time.Sleep(3 * time.Second)
	df.Cancel()
	if i != 3 {
		t.Fatalf("Error in Debounce")
	}
}
