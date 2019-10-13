package string

// PadStart pads string on the left side if it's shorter than length.
// Padding characters are truncated if they exceed length.
func PadStart(s string, l int, c string) string {
	var strLen int
	if l <= 0 {
		strLen = 0
	} else {
		strLen = len(s)
	}

	if strLen < l {
		return createPadding(l-strLen, c) + s
	}

	return s
}
