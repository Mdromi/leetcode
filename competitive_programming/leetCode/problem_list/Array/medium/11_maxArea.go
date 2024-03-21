package medium

/**
Problem: https://leetcode.com/problems/container-with-most-water/
*/

/*
Time: O(n)
Space: O(1)
*/

func MaxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	maxWater := 0

	for left < right {
		h := min(height[left], height[right])
		w := right - left
		maxWater = max(maxWater, h*w)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
