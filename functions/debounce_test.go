package functions

import (
	"sync"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {

	i := 0
	mux := sync.Mutex{}
	h := func() {
		mux.Lock()
		increment(&i)
		mux.Unlock()
	}
	df := Debounce(h, 1000)
	df.Call()
	time.Sleep(3 * time.Second)
	df.Cancel()
	mux.Lock()
	if i != 3 {
		t.Log(i)
		t.Fail()

	}
	mux.Unlock()
}
func increment(i *int) {
	*i += 1
}
