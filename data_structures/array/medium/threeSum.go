package medium

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates for the first number
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++ // Skip duplicates for the second number
				}
				for left < right && nums[right] == nums[right+1] {
					right-- // Skip duplicates for the third number
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
