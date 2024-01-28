package easy

/**
Problem: https://leetcode.com/problems/missing-number/
*/

/*
Time: O(n)
Space: O(1)
*/
func MissingNumber(nums []int) int {
	length := len(nums)
	s := length * (length + 1) / 2
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return s - sum
}
