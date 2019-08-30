package functions

import (
	"time"
)

// Debounced struct gives a debounced function which calls it after certain intervals
type Debounced struct {
	function func()
	duration int
	stop     bool
	channel  chan bool
}

// Debounce Function returns a functions for a passed function and a duration in milli-seconds
// On calling the returned function the passed function is called repeatedly after fixed intervals
// C
func Debounce(f func(), t int) Debounced {

	debounced := Debounced{
		f, t, false, make(chan bool),
	}
	return debounced
}

// Flush instantly invoke function
func (d *Debounced) Flush() {
	d.function()
}

// Call the Debounce Method so that the function repeatedly call itself
func (d *Debounced) Call() {
	d.stop = false
	go func() {

		for !d.stop {
			select {
			case hg := <-d.channel:
				d.stop = hg
			default:
				d.stop = false
			}
			go d.function()
			time.Sleep(time.Duration(int64(d.duration) * int64(time.Millisecond)))
		}
	}()
}

// Cancel stop the debounce method
func (d *Debounced) Cancel() {
	d.channel <- true
}
