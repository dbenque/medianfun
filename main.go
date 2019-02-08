package main

import (
	"fmt"
	"sync"
)

func main() {

	values := []int64{15, 2, 3, 99, 110, 79, 1, 6, 78, 1, 56, 2, 101, 5}

	injectChan := make(chan int64)
	medianChan := make(chan float64)

	medianCalculator := NewMedianCalculatorBasedOnHeap()
	medianCalculator.Init(len(values), medianChan)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for m := range medianChan {
			fmt.Printf("median: %f\n", m)
		}
		wg.Done()
	}()

	go medianCalculator.IngestInt(injectChan)

	for _, i := range values {
		injectChan <- i
	}
	close(injectChan)
	wg.Wait()
}
