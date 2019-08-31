package string

func PadEnd(s string, l int, c string) string {
	var strLen int
	if l <= 0 {
		strLen = 0
	} else {
		strLen = len(s)
	}

	if strLen < l {
		return s + createPadding(l-strLen, c)
	} else {
		return s
	}
}
