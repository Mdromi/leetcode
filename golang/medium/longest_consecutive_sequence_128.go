package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longestStreak := 0

	for num := range numSet {
		// Only check for the start of the sequence
		if _, found := numSet[num-1]; !found {
			currentNum := num
			currentStreak := 1

			// Check for consecutive numbers
			for {
				if _, found := numSet[currentNum+1]; found {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println("Longest consecutive sequence length (example 1):", longestConsecutive(nums1))
	fmt.Println("Longest consecutive sequence length (example 2):", longestConsecutive(nums2))
}
