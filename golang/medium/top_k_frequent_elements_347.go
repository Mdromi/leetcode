package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	num     int // the number itself
	freq    int // frequency of the number
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: Create a frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Create a min-heap
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Step 3: Insert elements into the min-heap
	for num, freq := range freqMap {
		heap.Push(minHeap, Element{num, freq})
		if minHeap.Len() > k {
			heap.Pop(minHeap) // Remove the smallest element (minimum frequency)
		}
	}

	// Step 4: Extract elements from the min-heap to result
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(minHeap).(Element).num
	}

	return result
}

func main() {
	// Example usage
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	fmt.Printf("Input: nums = %v, k = %d\n", nums1, k1)
	fmt.Println("Output:", topKFrequent(nums1, k1)) // Output: [1 2]

	nums2 := []int{1}
	k2 := 1
	fmt.Printf("Input: nums = %v, k = %d\n", nums2, k2)
	fmt.Println("Output:", topKFrequent(nums2, k2)) // Output: [1]
}
