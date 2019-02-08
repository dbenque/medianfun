package main

import "math"

type MedianCalculator interface {
	Init(count int, medianchan chan<- float64)
	IngestInt(intchan <-chan int64)
}
type MedianCalculatorBasedOnHeap struct {
	lowerHalf  *Heap
	upperHalf  *Heap
	medianchan chan<- float64
}

func NewMedianCalculatorBasedOnHeap() MedianCalculator {
	return &MedianCalculatorBasedOnHeap{}
}

func (m *MedianCalculatorBasedOnHeap) Init(count int, medianchan chan<- float64) {
	m.lowerHalf = NewHeapMax(count/2 + 1)
	m.upperHalf = NewHeapMin(count/2 + 1)
	m.medianchan = medianchan
}
func (m *MedianCalculatorBasedOnHeap) IngestInt(intchan <-chan int64) {

	GetMedian := func() {
		if m.lowerHalf.Size == m.upperHalf.Size {
			m.medianchan <- (float64(m.lowerHalf.Peek()) + float64(m.upperHalf.Peek())) / 2.0
			return
		}

		if m.lowerHalf.Size > m.upperHalf.Size {
			m.medianchan <- float64(m.lowerHalf.Peek())
			return
		}
		m.medianchan <- float64(m.upperHalf.Peek())
		return
	}

	for i := range intchan {
		if m.lowerHalf.Size == 0 {
			m.lowerHalf.Insert(i)
			GetMedian()
		} else {
			maxOfMinHalf := m.lowerHalf.Peek()
			if i > maxOfMinHalf {
				m.upperHalf.Insert(i)
			} else {
				m.lowerHalf.Insert(i)
			}
			if math.Abs(float64(m.upperHalf.Size-m.lowerHalf.Size)) <= 1 {
				GetMedian()
				continue
			}

			if m.upperHalf.Size > m.lowerHalf.Size {
				m.lowerHalf.Insert(m.upperHalf.Peek())
				m.upperHalf.Remove()
			} else {
				m.upperHalf.Insert(m.lowerHalf.Peek())
				m.lowerHalf.Remove()
			}
			GetMedian()
		}
	}
	close(m.medianchan)
}
