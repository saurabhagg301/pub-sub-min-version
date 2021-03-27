/*
2 packages - producer and consumer
art of module - pubsub
producer - generating stream of number - one number at a time
subcriber/consumer -  has to tell whether the number is odd or even
*/

package main

import (
	"github.com/saurabhagg301/pub-sub-min-version/pkg/consumer"
	"github.com/saurabhagg301/pub-sub-min-version/pkg/producer"
	"sync"
)

// global variable
var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go producer.GenerateRandomNumbers(ch, &wg)
	go consumer.ReadNumber(ch, &wg)

	wg.Add(2)

	wg.Wait()
}
