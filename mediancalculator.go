package main

type MedianCalculator interface {
	Init(count int, medianchan chan<- float64)
	IngestInt(intchan <-chan int64)
}
