package string

// Checks if string ends with the given target string. If position specified is <= 0,
// it will check until the end of the string
func EndsWith(s, t string, pos int64) bool {
	l := int64(len(s))

	if pos <= 0 {
		pos = l
	}

	if pos > l {
		pos = l
	}

	var end = pos
	pos -= int64(len(t))
	return pos >= 0 && s[pos:end] == t
}
