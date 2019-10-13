package slice

import (
	"math"
	"reflect"
)

func Chunk(sl interface{}, size int) []interface{} {
	slice, success := takeArg(sl, reflect.Slice)
	if !success {
		panic("You must pass in a slice value!")
	}

	// return empty slice if size or length of slice is 0
	l := slice.Len()
	if l <= 0 || size < 1 {
		return make([]interface{}, 0)
	}

	index := 0
	resIndex := 0

	result := make([]interface{}, int(math.Ceil(float64(l)/float64(size))))

	for index < l {
		if index+size < slice.Len() {
			result[resIndex] = slice.Slice(index, index+size).Interface()
		} else {
			result[resIndex] = slice.Slice(index, slice.Len()).Interface()
		}

		resIndex++
		index += size
	}
	return result
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
