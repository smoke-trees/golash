package number

import "reflect"

func Clamp(n, l, u interface{}) interface{} {
	if reflect.TypeOf(n) != reflect.TypeOf(l) || reflect.TypeOf(n) != reflect.TypeOf(u) {
		panic("All arguments must be of the same type!")
	}
	switch n.(type) {
	case uint8:
		if n.(uint8) > u.(uint8) {
			n = u
		}

		if n.(uint8) < l.(uint8) {
			n = l
		}
		return n.(uint8)
	case int8:
		if n.(int8) > u.(int8) {
			n = u
		}

		if n.(int8) < l.(int8) {
			n = l
		}
		return n.(int8)
	case uint16:
		if n.(uint16) > u.(uint16) {
			n = u
		}

		if n.(uint16) < l.(uint16) {
			n = l
		}
		return n.(uint16)
	case int16:
		if n.(int16) > u.(int16) {
			n = u
		}

		if n.(int16) < l.(int16) {
			n = l
		}
		return n.(int16)
	case uint32:
		if n.(uint32) > u.(uint32) {
			n = u
		}

		if n.(uint32) < l.(uint32) {
			n = l
		}
		return n.(uint32)
	case int32:
		if n.(int32) > u.(int32) {
			n = u
		}

		if n.(int32) < l.(int32) {
			n = l
		}
		return n.(int32)
	case uint64:
		if n.(uint64) > u.(uint64) {
			n = u
		}

		if n.(uint64) < l.(uint64) {
			n = l
		}
		return n.(uint64)
	case int64:
		if n.(int64) > u.(int64) {
			n = u
		}

		if n.(int64) < l.(int64) {
			n = l
		}
		return n.(int64)
	case int:
		if n.(int) > u.(int) {
			n = u
		}

		if n.(int) < l.(int) {
			n = l
		}
		return n.(int)
	case float32:
		if n.(float32) > u.(float32) {
			n = u
		}

		if n.(float32) < l.(float32) {
			n = l
		}
		return n.(float32)
	case float64:
		if n.(float64) > u.(float64) {
			n = u
		}

		if n.(float64) < l.(float64) {
			n = l
		}
		return n.(float64)
	}
	panic("Invalid type! The Clamp function accepts only numeric types.")
}
