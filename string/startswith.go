package string

func StartsWith(s, t string, p int) bool {
	strLen := len(s)
	if p < 0 {
		p = 0
	} else if p > strLen {
		p = strLen
	}

	defer func() {
		recover()
	}()

	return s[p:p+len(t)] == t
}
