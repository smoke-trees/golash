package function

import (
	"testing"
	"time"
)

// TestDelay we can use Non-Blocking channels
// If after a wait time there is no buffer to read from
// the channel then the test fails
func TestDelay(t *testing.T) {
	ch := make(chan string)

	f := func() {
		ch <- "hi"
	}
	Delay(f, 3000)
	time.Sleep(3 * time.Second)
	select {
	case h := <-ch:
		{
			if h != "hi" {
				t.Logf("Recieved :%s", h)
				t.Fail()
			}
		}
	default:
		t.Fail()
	}
}
