package producer

import (
	"sync"
)

// GenerateRandomNumbers to generate numbers
func GenerateRandomNumbers(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	// fmt.Println(">>> Debug: Producer package: Decrementing wg counter <<<<")
	wg.Done()
}
