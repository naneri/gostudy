package circuit_breaker

import (
	"errors"
	"sync"
	"time"
)

type CircuitBreaker struct {
	// maxFailures is the number of consecutive failures allowed before tripping the circuit breaker
	maxFailures int
	// timeout is the duration for which to open the circuit breaker after tripping it
	timeout time.Duration

	state               int // 0 - closed, 1 - open, 2 - half-open
	consecutiveFailures int
	lastFailure         time.Time

	mutex sync.Mutex
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       0,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	// If the circuit breaker is in the open state, check if it's time to switch to half-open
	if cb.state == 1 && time.Now().Sub(cb.lastFailure) >= cb.timeout {
		cb.state = 2
	}

	// If the circuit breaker is in the half-open state, try the function and see if it succeeds
	if cb.state == 2 {
		err := fn()
		if err == nil {
			cb.consecutiveFailures = 0
			cb.state = 0
			return nil
		} else {
			cb.consecutiveFailures++
			if cb.consecutiveFailures >= cb.maxFailures {
				cb.state = 1
				cb.lastFailure = time.Now()
			}
			return err
		}
	}

	// If the circuit breaker is in the closed state, try the function and see if it succeeds
	if cb.state == 0 {
		err := fn()
		if err == nil {
			cb.consecutiveFailures = 0
			return nil
		} else {
			cb.consecutiveFailures++
			if cb.consecutiveFailures >= cb.maxFailures {
				cb.state = 1
				cb.lastFailure = time.Now()
			}
			return err
		}
	}

	// If the circuit breaker is in the open state, return an error immediately
	return errors.New("circuit breaker is open")
}
