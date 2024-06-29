package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Initialize the answer array with 1s
	answer := make([]int, n)
	answer[0] = 1

	// Step 1: Calculate left products
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// Step 2: Calculate right products and update the answer array
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right *= nums[i]
	}

	return answer
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Println("Output:", productExceptSelf(nums1)) // Output: [24, 12, 8, 6]

	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Println("Output:", productExceptSelf(nums2)) // Output: [0, 0, 9, 0, 0]
}
