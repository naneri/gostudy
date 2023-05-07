package main

import (
	"fmt"
	"time"
)

func main() {

}

func endingFunc(closeTimer <-chan time.Time) {
	for {
		select {
		case <-closeTimer:
			fmt.Println("ended")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
