package easy

/**
Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
*/

/*
Time: O(n)
Space: O(1)
*/
func RemoveElement(nums []int, val int) int {
	// Initialize a pointer 'i' to track the position for overwriting non-target values
	i := 0

	// Iterate through each element in the array
	for _, num := range nums {
		// Check if the current element is not equal to the target value
		if num != val {
			// Overwrite the original array at position 'i' with the non-target value
			nums[i] = num
			// Increment the pointer 'i' to the next position
			i++
		}
	}

	// 'i' represents the new length of the array without the specified value
	return i
}
