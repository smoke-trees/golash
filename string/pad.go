package string

import "math"

// Pad pads string on the left and right sides if it's shorter than length.
// Padding characters are truncated if they can't be evenly divided by length.
func Pad(s string, l int, c string) string {
	var strLen int
	if l > 0 {
		strLen = len(s)
	} else {
		strLen = 0
	}

	if l == 0 || strLen >= l {
		return s
	}

	mid := float64(l-strLen) / 2

	return createPadding(int(math.Floor(mid)), c) + s + createPadding(int(math.Ceil(mid)), c)
}

func createPadding(l int, c string) string {
	strLen := len(c)
	if strLen < 2 {
		if strLen != 0 {
			return Repeat(c, int64(l))
		}

		return c
	}

	result := Repeat(c, int64(math.Ceil(float64(l)/float64(strLen))))
	return result[:l]
}
