package main

import (
	"testing"
)

func TestParentIndex(t *testing.T) {

	tests := []struct {
		name   string
		index  int
		parent int
	}{
		{
			name:   "zero",
			index:  0,
			parent: 0,
		},
		{
			name:   "1",
			index:  1,
			parent: 0,
		},
		{
			name:   "2",
			index:  2,
			parent: 0,
		},
		{
			name:   "3",
			index:  3,
			parent: 1,
		},
		{
			name:   "4",
			index:  4,
			parent: 1,
		},
		{
			name:   "5",
			index:  5,
			parent: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParentIndex(tt.index); got != tt.parent {
				t.Errorf("ParentIndex() = %v, want %v", got, tt.parent)
			}
		})
	}
}

func TestLeftChild(t *testing.T) {

	tests := []struct {
		name      string
		parent    int
		leftChild int
	}{
		{
			name:      "zero",
			parent:    0,
			leftChild: 1,
		},
		{
			name:      "1",
			parent:    1,
			leftChild: 3,
		},
		{
			name:      "2",
			parent:    2,
			leftChild: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftChild(tt.parent); got != tt.leftChild {
				t.Errorf("LeftChild() = %v, want %v", got, tt.leftChild)
			}
		})
	}
}
