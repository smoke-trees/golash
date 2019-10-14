package function

// Counter use Counter.Call() method to Run the passed function in Times 1 time
type Counter struct {
	f     func(int)
	n     int
	count int
}

// Times Returns a Counter struct with a Call Method to invoke Calling of the method
func Times(f func(int), n int) Counter {
	c := Counter{
		f:     f,
		n:     n,
		count: 0,
	}
	return c
}

// Call the function in time 1 time
func (c *Counter) Call() {
	go func() {
		for c.count < c.n {
			c.f(c.count)
			c.count++
		}
	}()
}
