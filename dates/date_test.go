package dates

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	n:=Now()
	gn := time.Now().Unix()
	if n != gn {
		t.Fatalf("Error in Now function: n:%d gn:%d",n,gn)
	}
}
