package easy

/**
Problem: https://leetcode.com/problems/search-insert-position/
*/

// searchInsert searches for the target value in a sorted array of integers.
// If the target is found, it returns the index of the target.
// If the target is not found, it returns the index where the target would be inserted to maintain sorted order.
// The function uses a binary search algorithm to achieve O(log n) runtime complexity.
func SearchInsert(nums []int, target int) int {
	// Initialize left and right pointers for binary search.
	left := 0
	right := len(nums) - 1

	// Perform binary search until the left pointer is greater than the right pointer.
	for left <= right {
		// Calculate the middle index of the current search space.
		mid := (left + right) / 2
		midValue := nums[mid] // Retrieve the value at the middle index.

		// Check if the middle value is equal to the target.
		if midValue == target {
			return mid // Return the index of the target.
		} else if midValue < target {
			left = mid + 1 // Adjust the search space to the right half.
		} else {
			right = mid - 1 // Adjust the search space to the left half.
		}
	}

	// If the target is not found, return the left pointer, which represents the index where the target would be inserted.
	return left
}
