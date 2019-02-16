package main

import (
	"math/rand"
	"sync"
	"testing"
)

var values [10000]int64

func init() {
	for i := 0; i < 10000; i++ {
		values[i] = rand.Int63()
	}

}
func BenchmarkHeap(b *testing.B) {

	// run the function b.N times
	for n := 0; n < b.N; n++ {

		injectChan := make(chan int64)
		medianChan := make(chan float64)

		medianCalculator := NewMedianCalculatorBasedOnHeap()
		medianCalculator.Init(len(values), medianChan)

		wg := sync.WaitGroup{}
		wg.Add(1)
		//var final float64
		go func() {
			for m := range medianChan {
				//fmt.Printf("median: %f\n", m)
				final := m
				final++
			}
			wg.Done()
		}()

		go medianCalculator.IngestInt(injectChan)

		for _, i := range values {
			injectChan <- i
		}
		close(injectChan)
		wg.Wait()

		//fmt.Printf("median: %f\n", final)

	}
}

func BenchmarkListMed(b *testing.B) {

	// run the function b.N times
	for n := 0; n < b.N; n++ {
		injectChan := make(chan int64)
		medianChan := make(chan float64)

		medianCalculator := NewDoubleListSolution()
		medianCalculator.Init(len(values), medianChan)

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for m := range medianChan {
				//fmt.Printf("median: %f\n", m)
				final := m
				final++
			}
			wg.Done()
		}()

		go medianCalculator.IngestInt(injectChan)

		for _, i := range values {
			injectChan <- i
		}
		close(injectChan)
		wg.Wait()

		// /fmt.Printf("median: %f\n", final)
	}
}
