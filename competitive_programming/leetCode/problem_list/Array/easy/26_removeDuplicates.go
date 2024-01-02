package easy

/**
Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
*/

/*
Time: O(n)
Space: O(1)
*/

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[uniqueCount] = nums[i]
			uniqueCount++
		}
	}

	return uniqueCount
}
