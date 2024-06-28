package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// Step 1: Create a frequency map
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Create a bucket to store elements by frequency
	maxFreq := len(nums) // maximum frequency can be nums.length
	bucket := make([][]int, maxFreq+1)
	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	// Step 3: Collect k most frequent elements
	result := make([]int, 0, k)
	for freq := maxFreq; freq > 0 && len(result) < k; freq-- {
		if len(bucket[freq]) > 0 {
			result = append(result, bucket[freq]...)
		}
	}

	return result[:k] // return exactly k elements
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
