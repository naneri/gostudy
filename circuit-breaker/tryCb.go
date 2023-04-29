package circuit_breaker

import (
	"errors"
	"fmt"
	"time"
)

func TryCb() {
	cb := NewCircuitBreaker(3, 5*time.Second)

	// Call a function that will succeed
	err := cb.Execute(func() error {
		fmt.Println("Hello, world!")
		return nil
	})
	fmt.Println(err)

	for i := 0; i < 5; i++ {
		func() {
			// Call the same function again to trigger the circuit breaker
			err = cb.Execute(func() error {
				return errors.New("something went wrong")
			})
			fmt.Println(err)
		}()
	}

	// Wait for the timeout to elapse
	time.Sleep(6 * time.Second)

	err = cb.Execute(func() error {
		return errors.New("something went wrong")
	})
	fmt.Println(err)

	// Call the function again to see if the circuit breaker has switched to half-open
	err = cb.Execute(func() error {
		fmt.Println("Hello, world!")
		return nil
	})
	fmt.Println(err)
}
