package string

// PadEnd pads string on the right side if it's shorter than length.
// Padding characters are truncated if they exceed length.
func PadEnd(s string, l int, c string) string {
	var strLen int
	if l <= 0 {
		strLen = 0
	} else {
		strLen = len(s)
	}

	if strLen < l {
		return s + createPadding(l-strLen, c)
	}

	return s
}
