package consumer

import (
	"fmt"
	"sync"
)

// ReadNumber to read number from channel ch
func ReadNumber(ch chan int, wg *sync.WaitGroup) {
	flagChannelClosed := false
	for flagChannelClosed == false {
		select {
		case n, ok := <-ch:
			if ok {
				if n%2 == 0 {
					fmt.Println("Number = ", n, ". It's a even number")
				} else {
					fmt.Println("Number = ", n, ". It's a odd number")
				}
			} else {
				flagChannelClosed = true
				// channel closed
				// fmt.Println(">>> Debug: Consumer package: Decrementing wg counter <<<<")
				wg.Done()
			}
		default:
			//fmt.Println("waiting for number..")

		}
	}
	// for {
	// 	n := <-ch
	// 	if n%2 == 0 {
	// 		fmt.Println("Number = ", n, ". It's a even number")
	// 	} else {
	// 		fmt.Println("Number = ", n, ". It's a odd number")
	// 	}
	// 	// time.Sleep(1 * time.Millisecond)
	// }

}
