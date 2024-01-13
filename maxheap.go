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

// maxheapifyUp will  heapify from bottom top
func (h *MaxHeap) maxheapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
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

	buildHeap := []int{10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

}
