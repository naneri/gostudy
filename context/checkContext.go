package context

import (
	"context"
	"fmt"
	"time"
)

func TryOutContext() {
	newCtx, cancel := context.WithCancel(context.Background())
	go someFuncWithContext(newCtx)
	go AnotherFuncWIthContext(newCtx)
	time.Sleep(3 * time.Second)
	fmt.Println("ctx cancelled")
	cancel()
	time.Sleep(5 * time.Second)

}
func AnotherFuncWIthContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("we are fone here")
			return
		default:
			time.Sleep(5 * time.Second)
			fmt.Println("here have waited 5 seconds")
		}
	}
}

func someFuncWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("we are done, bye bye")
			return
		default:
			fmt.Println("hi here")
			time.Sleep(1 * time.Second)
		}
	}

}
