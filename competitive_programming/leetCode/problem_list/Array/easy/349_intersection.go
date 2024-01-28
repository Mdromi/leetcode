package easy

/**
Problem: https://leetcode.com/problems/intersection-of-two-arrays/
*/

/*
Time: O(m + n)
Space: O(m)
*/

func Intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, num := range nums1 {
		set1[num] = true
	}

	for _, num := range nums2 {
		set2[num] = true
	}

	var result []int
	for num := range set1 {
		if set2[num] {
			result = append(result, num)
		}
	}

	return result
}
