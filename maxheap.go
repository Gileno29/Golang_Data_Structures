package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxheapifyUp(len(h.array) - 1)
}

// Extract return the largest key, and 	removes it from the heap

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Cannot extract becaseu arrya length is 0")
		return -1
	}

	// take out the last index
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxheapifyDown(0)

	return extracted
}

// maxheapifyUp will  heapify from bottom top
func (h *MaxHeap) maxheapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxheapfyDown  will heapify top to bottom

func (h *MaxHeap) maxheapifyDown(index int) {
	// loop while  index has at last one child

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the child index
func left(i int) int {
	return 2*i + 1
}

// get the rigth child index
func right(i int) int {
	return 2*1 + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main() {
	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 5, 7, 11, 12, 22, 30, 31, 33}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
