package main

import "fmt"

type Node struct {
	value    int64
	next     *Node
	previous *Node
}

type DoubleListMedian struct {
	head   *Node
	tail   *Node
	median *Node
	odd    bool
}

func NewDoubleListMedian() *DoubleListMedian {
	return &DoubleListMedian{
		odd: false,
	}
}

func (d *DoubleListMedian) Print() {
	n := d.head

	for n != nil {
		if n == d.median {
			fmt.Printf("#")
		}
		fmt.Printf("%d ", n.value)
		n = n.next
	}
	fmt.Println()
}

func (d *DoubleListMedian) Insert(value int64) {
	defer func() { d.odd = !d.odd }()
	n := &Node{value: value}
	if d.head == nil {
		d.head = n
		d.tail = n
		d.median = n
		return
	}

	if d.median.value <= value {

		if !d.odd {
			defer func() {
				d.median = d.median.next
			}()
		}

		next := d.median.next
		for next != nil && next.value <= value {
			next = next.next
		}

		if next == nil {
			d.tail.next = n
			n.previous = d.tail
			d.tail = n
		} else {
			n.next = next
			n.previous = next.previous
			n.previous.next = n
			next.previous = n
		}
	} else {

		if d.odd {
			defer func() {
				d.median = d.median.previous
			}()
		}

		previous := d.median.previous
		for previous != nil && previous.value >= value {
			previous = previous.previous
		}

		if previous == nil {
			d.head.previous = n
			n.next = d.head
			d.head = n
		} else {
			n.previous = previous
			n.next = previous.next
			previous.next = n
			n.next.previous = n
		}
	}

}
