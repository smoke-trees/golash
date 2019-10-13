package number

import (
	"math"
	"reflect"
)


// InRange checks whether e is present in between n to s
// Supported Types : int8, uint8, uint16, int16, uint32, int32, uint64, int64, int, uint, float32, float64
func InRange(n, s, e interface{}) bool {
	if reflect.TypeOf(n) != reflect.TypeOf(s) || reflect.TypeOf(n) != reflect.TypeOf(e) {
		panic("All arguments must be of the same type!")
	}

	switch n.(type) {
	case uint8:
		return n.(uint8) >= uint8(math.Min(float64(s.(uint8)), float64(e.(uint8)))) && n.(uint8) < uint8(math.Max(float64(s.(uint8)), float64(e.(uint8))))
	case int8:
		return n.(int8) >= int8(math.Min(float64(s.(int8)), float64(e.(int8)))) && n.(int8) < int8(math.Max(float64(s.(int8)), float64(e.(int8))))
	case uint16:
		return n.(uint16) >= uint16(math.Min(float64(s.(uint16)), float64(e.(uint16)))) && n.(uint16) < uint16(math.Max(float64(s.(uint16)), float64(e.(uint16))))
	case int16:
		return n.(int16) >= int16(math.Min(float64(s.(int16)), float64(e.(int16)))) && n.(int16) < int16(math.Max(float64(s.(int16)), float64(e.(int16))))
	case uint32:
		return n.(uint32) >= uint32(math.Min(float64(s.(uint32)), float64(e.(uint32)))) && n.(uint32) < uint32(math.Max(float64(s.(uint32)), float64(e.(uint32))))
	case int32:
		return n.(int32) >= int32(math.Min(float64(s.(int32)), float64(e.(int32)))) && n.(int32) < int32(math.Max(float64(s.(int32)), float64(e.(int32))))
	case uint64:
		return n.(uint64) >= uint64(math.Min(float64(s.(uint64)), float64(e.(uint64)))) && n.(uint64) < uint64(math.Max(float64(s.(uint64)), float64(e.(uint64))))
	case int64:
		return n.(int64) >= int64(math.Min(float64(s.(int64)), float64(e.(int64)))) && n.(int64) < int64(math.Max(float64(s.(int64)), float64(e.(int64))))
	case uint:
		return n.(uint) >= uint(math.Min(float64(s.(uint)), float64(e.(uint)))) && n.(uint) < uint(math.Max(float64(s.(uint)), float64(e.(uint))))
	case int:
		return n.(int) >= int(math.Min(float64(s.(int)), float64(e.(int)))) && n.(int) < int(math.Max(float64(s.(int)), float64(e.(int))))
	case float32:
		return n.(float32) >= float32(math.Min(float64(s.(float32)), float64(e.(float32)))) && n.(float32) < float32(math.Max(float64(s.(float32)), float64(e.(float32))))
	case float64:
		return n.(float64) >= math.Min(s.(float64), e.(float64)) && n.(float64) < math.Max(s.(float64), e.(float64))
	}

	panic("Invalid type! The Clamp function accepts only numeric types.")
}
