package string

import "math"

// Repeat repeats the given string n times.
func Repeat(s string, n int64) string {
	result := ""
	if n < 1 || n > math.MaxInt64 {
		return result
	}

	for {
		if n%2 != 0 {
			result += s
		}
		n = int64(math.Floor(float64(n) / 2))
		if n != 0 {
			s += s
		}

		if n == 0 {
			break
		}
	}
	return result
}
