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
		return inRangeUint8(n, s, e)
	case int8:
		return inRangeInt8(n, s, e)
	case uint16:
		return inRangeUint16(n, s, e)
	case int16:
		return inRangeInt16(n, s, e)
	case uint32:
		return inRangeUint32(n, s, e)
	case int32:
		return inRangeInt32(n, s, e)
	case uint64:
		return inRangeUint64(n, s, e)
	case int64:
		return inRangeInt64(n, s, e)
	case uint:
		return inRangeUint(n, s, e)
	case int:
		return inRangeInt(n, s, e)
	case float32:
		return inRangeFloat32(n, s, e)
	case float64:
		return inRangeFloat64(n, s, e)
	}

	panic("Invalid type! The Clamp function accepts only numeric types.")
}

func inRangeFloat64(n interface{}, s interface{}, e interface{}) bool {
	return n.(float64) >= math.Min(s.(float64), e.(float64)) && n.(float64) < math.Max(s.(float64), e.(float64))
}

func inRangeFloat32(n interface{}, s interface{}, e interface{}) bool {
	return n.(float32) >= float32(math.Min(float64(s.(float32)), float64(e.(float32)))) && n.(float32) < float32(math.Max(float64(s.(float32)), float64(e.(float32))))
}

func inRangeInt(n interface{}, s interface{}, e interface{}) bool {
	return n.(int) >= int(math.Min(float64(s.(int)), float64(e.(int)))) && n.(int) < int(math.Max(float64(s.(int)), float64(e.(int))))
}

func inRangeInt64(n interface{}, s interface{}, e interface{}) bool {
	return n.(int64) >= int64(math.Min(float64(s.(int64)), float64(e.(int64)))) && n.(int64) < int64(math.Max(float64(s.(int64)), float64(e.(int64))))
}

func inRangeUint(n interface{}, s interface{}, e interface{}) bool {
	return n.(uint) >= uint(math.Min(float64(s.(uint)), float64(e.(uint)))) && n.(uint) < uint(math.Max(float64(s.(uint)), float64(e.(uint))))
}

func inRangeUint64(n interface{}, s interface{}, e interface{}) bool {
	return n.(uint64) >= uint64(math.Min(float64(s.(uint64)), float64(e.(uint64)))) && n.(uint64) < uint64(math.Max(float64(s.(uint64)), float64(e.(uint64))))
}

func inRangeInt32(n interface{}, s interface{}, e interface{}) bool {
	return n.(int32) >= int32(math.Min(float64(s.(int32)), float64(e.(int32)))) && n.(int32) < int32(math.Max(float64(s.(int32)), float64(e.(int32))))
}

func inRangeUint32(n interface{}, s interface{}, e interface{}) bool {
	return n.(uint32) >= uint32(math.Min(float64(s.(uint32)), float64(e.(uint32)))) && n.(uint32) < uint32(math.Max(float64(s.(uint32)), float64(e.(uint32))))
}

func inRangeInt16(n interface{}, s interface{}, e interface{}) bool {
	return n.(int16) >= int16(math.Min(float64(s.(int16)), float64(e.(int16)))) && n.(int16) < int16(math.Max(float64(s.(int16)), float64(e.(int16))))
}

func inRangeUint16(n interface{}, s interface{}, e interface{}) bool {
	return n.(uint16) >= uint16(math.Min(float64(s.(uint16)), float64(e.(uint16)))) && n.(uint16) < uint16(math.Max(float64(s.(uint16)), float64(e.(uint16))))
}

func inRangeInt8(n interface{}, s interface{}, e interface{}) bool {
	return n.(int8) >= int8(math.Min(float64(s.(int8)), float64(e.(int8)))) && n.(int8) < int8(math.Max(float64(s.(int8)), float64(e.(int8))))
}

func inRangeUint8(n interface{}, s interface{}, e interface{}) bool {
	return n.(uint8) >= uint8(math.Min(float64(s.(uint8)), float64(e.(uint8)))) && n.(uint8) < uint8(math.Max(float64(s.(uint8)), float64(e.(uint8))))
}
