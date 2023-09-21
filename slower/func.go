package slower

import "time"

// SlowingFunc starts to slow down twice after every iteration in case it does not succeed.
// It does not handle errors.
func SlowingFunc(funcToRun func() error) {
	delay := time.Second
	threshold := time.Minute
	for {
		err := funcToRun()
		if err == nil {
			return
		}
		time.Sleep(delay)
		if delay < threshold {
			delay *= 2
		}
	}
}
