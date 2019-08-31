package functions

import "time"

// Delay invokes the function f after w milliseconds
func Delay(f func(), w int64) {
	if w < 100 {
		w = 100
	}
	go func() {
		time.Sleep(time.Duration(w * int64(time.Millisecond)))
		f()
	}()
}
