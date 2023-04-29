package circuit_breaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloud-native-go/examples/ch04"
	"math/rand"
	"time"
)

func TryCnCb() {
	// Create a circuit breaker that allows up to 3 consecutive failures
	breaker := ch04.Breaker(failingService, 3)

	// Make some requests to the service through the circuit breaker
	for i := 0; i < 20; i++ {
		result, err := breaker(context.Background())
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Result: %s\n", result)
		}

		// Sleep for a short time to simulate a delay between requests
		time.Sleep(100 * time.Millisecond)
	}
}

// Define a function that simulates a service that sometimes fails
func failingService(ctx context.Context) (string, error) {
	if rand.Intn(10) < 7 {
		return "", errors.New("service unavailable")
	}
	return "Hello, world!", nil
}
