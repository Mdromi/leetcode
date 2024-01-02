package easy

/**
Problem: https://leetcode.com/problems/two-sum/description/
*/

/*
Time: O(n*2)
Space: O(1)
*/
func TwoSumLinearSearch(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{} // Return an empty slice if no matching indices are found
}

/*
Time: O(n)
Space: O(n)
*/
func TwoSumHasMap(nums []int, target int) []int {
	// Create a map to store the indices of elements
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num

		// Check if the complement is already in the map
		if index, found := numMap[complement]; found {
			// Return the indices of the two numbers that add up to the target
			return []int{index, i}
		}

		// If not, store the current number and its index in the map
		numMap[num] = i
	}

	// If no solution is found
	return nil
}
