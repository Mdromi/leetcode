package medium

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if currentSum == target {
				return target
			}

			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum
			}

			if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
