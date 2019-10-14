package number

import "reflect"

// Clamp clamps number within the inclusive lower and upper bounds.
// Supported Types : int8, uint8, uint16, int16, uint32, int32, uint64, int64, int, uint, float32, float64
func Clamp(n, l, u interface{}) interface{} {
	if reflect.TypeOf(n) != reflect.TypeOf(l) || reflect.TypeOf(n) != reflect.TypeOf(u) {
		panic("All arguments must be of the same type!")
	}
	switch n.(type) {
	case uint8:
		return clampUint8(n, u, l)
	case int8:
		return clampInt8(n, u, l)
	case uint16:
		return clampUint16(n, u, l)
	case int16:
		return clampInt16(n, u, l)
	case uint32:
		return clampUint32(n, u, l)
	case int32:
		return clampInt32(n, u, l)
	case uint64:
		return clampUint64(n, u, l)
	case int64:
		return clampInt64(n, u, l)
	case uint:
		return clampUint(n, u, l)
	case int:
		return clampInt(n, u, l)
	case float32:
		return clampFloat32(n, u, l)
	case float64:
		return clampFloat64(n, u, l)
	}
	panic("Invalid type! The Clamp function accepts only numeric types.")
}

func clampUint8(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(uint8) > u.(uint8) {
		n = u
	}
	if n.(uint8) < l.(uint8) {
		n = l
	}
	return n.(uint8)
}

func clampInt8(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(int8) > u.(int8) {
		n = u
	}
	if n.(int8) < l.(int8) {
		n = l
	}
	return n.(int8)
}

func clampUint16(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(uint16) > u.(uint16) {
		n = u
	}
	if n.(uint16) < l.(uint16) {
		n = l
	}
	return n.(uint16)
}

func clampInt16(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(int16) > u.(int16) {
		n = u
	}
	if n.(int16) < l.(int16) {
		n = l
	}
	return n.(int16)
}

func clampUint32(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(uint32) > u.(uint32) {
		n = u
	}
	if n.(uint32) < l.(uint32) {
		n = l
	}
	return n.(uint32)
}

func clampInt32(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(int32) > u.(int32) {
		n = u
	}
	if n.(int32) < l.(int32) {
		n = l
	}
	return n.(int32)
}

func clampUint64(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(uint64) > u.(uint64) {
		n = u
	}
	if n.(uint64) < l.(uint64) {
		n = l
	}
	return n.(uint64)
}

func clampInt64(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(int64) > u.(int64) {
		n = u
	}
	if n.(int64) < l.(int64) {
		n = l
	}
	return n.(int64)
}

func clampUint(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(uint) > u.(uint) {
		n = u
	}
	if n.(uint) < l.(uint) {
		n = l
	}
	return n.(uint)
}

func clampInt(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(int) > u.(int) {
		n = u
	}
	if n.(int) < l.(int) {
		n = l
	}
	return n.(int)
}

func clampFloat32(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(float32) > u.(float32) {
		n = u
	}
	if n.(float32) < l.(float32) {
		n = l
	}
	return n.(float32)
}

func clampFloat64(n interface{}, u interface{}, l interface{}) interface{} {
	if n.(float64) > u.(float64) {
		n = u
	}
	if n.(float64) < l.(float64) {
		n = l
	}
	return n.(float64)
}
