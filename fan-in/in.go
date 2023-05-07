package fan_in

import "sync"

// FanIn function gets a slice of inputChannels, and returns and output Channel.
// The channel will be closed in this function after the input channels are depleted
func FanIn(inputChs ...chan int) chan int {
	outCh := make(chan int)

	go func() {
		ProcessAll(inputChs, outCh)
		close(outCh)
	}()

	return outCh
}

func ProcessAll(inputChs []chan int, outCh chan<- int) {
	wg := &sync.WaitGroup{}

	for _, inputCh := range inputChs {
		wg.Add(1)
		// this line would spawn the same number of goroutines as the number of input channels. That might not be optimal, depending on the number of inputChannels created. But the idea is, that the creator of the input channel slice will control himself the number of input channels created.
		go SendToOut(wg, inputCh, outCh)
	}

	wg.Wait()
}

func SendToOut(wg *sync.WaitGroup, inputCh chan int, outCh chan<- int) {
	defer wg.Done()
	for item := range inputCh {
		outCh <- item
	}
}
