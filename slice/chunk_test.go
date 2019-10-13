package slice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}

	for i, val := range Chunk(arr, 3) {
		test := [][]int{{0, 1, 2}, {3, 4, 5}}
		if !reflect.DeepEqual(val, test[i]) {
			t.Fatal("It should return chunked slices!")
		}
	}

	for i, val := range Chunk(arr, 4) {
		test := [][]int{{0, 1, 2, 3}, {4, 5}}
		if !reflect.DeepEqual(val, test[i]) {
			t.Fatal("It should return the last chunk as remaining elements!")
		}
	}

	if len(Chunk(make([]int, 0), 4)) != 0 {
		t.Fatal("It should return an empty slice!")
	}

	if len(Chunk(arr, 0)) != 0 {
		t.Fatal("It should return an empty slice!")
	}
}

func TestChunkPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("It should panic upon receiving a type that isn't a slice")
		}
	}()

	Chunk("random", 0)
}
