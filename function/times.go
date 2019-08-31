package function

type Counter struct {
	f     func(int)
	n     int
	count int
}

// Returns a Counter struct with a Call Method to invoke Calling of the method
func Times(f func(int), n int) Counter {
	c := Counter{
		f:     f,
		n:     n,
		count: 0,
	}
	return c
}

func (c *Counter) Call() {
	go func() {
		for c.count < c.n {
			c.f(c.count)
			c.count++
		}
	}()
}
