package main

type DoubleListSolution struct {
	dl         DoubleListMedian
	medianchan chan<- float64
}

func NewDoubleListSolution() MedianCalculator {
	return &DoubleListSolution{}
}

func (d *DoubleListSolution) Init(count int, medianchan chan<- float64) {
	d.medianchan = medianchan
}
func (d *DoubleListSolution) IngestInt(intchan <-chan int64) {

	GetMedian := func() {
		if d.dl.odd {
			d.medianchan <- float64(d.dl.median.value)
		} else {
			d.medianchan <- (float64(d.dl.median.value) + float64(d.dl.median.next.value)) / 2.0
		}
	}

	for i := range intchan {
		d.dl.Insert(i)
		GetMedian()
	}
	close(d.medianchan)
}
