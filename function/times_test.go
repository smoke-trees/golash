package function

import "testing"

func TestTimes(t *testing.T) {
	ch := make(chan int)
	f := func(n int) {
		ch <- n
	}
	c := Times(f, 3)
	c.Call()
	a := <-ch
	b := <-ch
	d := <-ch
	if a != 0 || b != 1 || d != 2 {
		t.Logf("a: %d | b: %d | d: %d\nExpected\na: %d | b: %d | d: %d\n", a, b, d, 0, 1, 2)
		t.Fail()
	}
}
