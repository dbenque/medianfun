package main

import "fmt"

type Heap struct {
	Size    int
	Array   []int64
	Compare func(i, j int64) bool
}

func NewHeapMin(maxSize int) *Heap {
	return &Heap{
		Array: make([]int64, maxSize),
		Compare: func(i, j int64) bool {
			return i < j
		},
	}
}
func NewHeapMax(maxSize int) *Heap {
	return &Heap{
		Array: make([]int64, maxSize),
		Compare: func(i, j int64) bool {
			return i > j
		},
	}
}

func (h *Heap) Remove() error {
	if h.Size == 0 {
		return nil
	}
	h.Array[0] = h.Array[h.Size-1]
	h.Size--
	if h.Size == 0 {
		return nil
	}

	current := 0
	for {

		childIndex := LeftChild(current)
		if childIndex > h.Size-1 {
			//already a leaf
			return nil
		}

		if childIndex < h.Size-1 {
			if !h.Compare(h.Array[childIndex], h.Array[childIndex+1]) {
				childIndex++ // take right
			}
		}

		if !h.Compare(h.Array[current], h.Array[childIndex]) {
			h.Array[current], h.Array[childIndex] = h.Array[childIndex], h.Array[current]
			current = childIndex
			continue
		}
		break
	}
	return nil
}

func (h *Heap) Insert(i int64) error {
	//TODO check N
	if h.Size >= len(h.Array) {
		return fmt.Errorf("Maxsize reached")
	}

	defer h.pushUp()

	h.Array[h.Size] = i
	h.Size++
	return nil
}

func (h *Heap) Peek() int64 {
	return h.Array[0]
}

func ParentIndex(n int) int {
	if n <= 2 {
		return 0
	}
	return (n - 1) / 2
}

func LeftChild(n int) int {
	return 2*n + 1
}

func (h *Heap) pushUp() {
	current := h.Size - 1
	parentIndex := ParentIndex(current)
	for h.Compare(h.Array[current], h.Array[parentIndex]) && current != 0 {
		h.Array[current], h.Array[parentIndex] = h.Array[parentIndex], h.Array[current]
		current = parentIndex
		parentIndex = ParentIndex(current)
	}
}

func (h *Heap) Print() {
	l := 1
	for i, j := range h.Array {
		fmt.Printf("%d ", j)
		if (i+2)%l == 0 {
			fmt.Printf("\n")
			l = l * 2
		}
	}
}
