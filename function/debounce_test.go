package function

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
	df := Debounce(h, 1*time.Second)
	df.Call()
	time.Sleep(3 * time.Second)
	df.Cancel()
	df.Flush()
	mux.Lock()
	if i == 4 || i == 5 {
		mux.Unlock()
		return
	}
	t.Log(i)
	t.Fail()
	mux.Unlock()
}
func increment(i *int) {
	*i++
}

func TestDebouncePanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fail()
		}
	}()

	_ = Debounce(func() {}, 1*time.Nanosecond)
}
