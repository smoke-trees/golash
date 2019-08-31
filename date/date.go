package date

import "time"

// Now returns the timestamp as number of seconds that have elapsed since the UNIX epoch
func Now() int64 {
	return time.Now().Unix()
}
