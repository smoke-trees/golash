package string

func PadStart(s string, l int, c string) string {
	var strLen int
	if l <= 0 {
		strLen = 0
	} else {
		strLen = len(s)
	}

	if strLen < l {
		return createPadding(l-strLen, c) + s
	} else {
		return s
	}
}
