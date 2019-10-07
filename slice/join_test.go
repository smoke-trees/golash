package slice

import (
	"testing"
)

func TestJoin(t *testing.T) {
	array := []string{"a", "b", "c"}

	if Join(array, "~") != "a~b~c" {
		t.Fatal("It should return join all array elements into a string")
	}
}

func TestJoinPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("It should panic upon receiving a type that isn't a slice")
		}
	}()

	Join("a", "b")
}
