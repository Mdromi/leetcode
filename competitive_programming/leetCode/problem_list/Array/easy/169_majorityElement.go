package easy

/**
Problem: https://leetcode.com/problems/majority-element/
*/

/*
Time: O(n)
Space: O(n)
*/

func MajorityElement_usingHashMap(nums []int) int {
	// Create a hash map to store the frequency of each element
	frequencyMap := make(map[int]int)

	// Iterate through the array to count the frequency of each element
	for _, num := range nums {
		frequencyMap[num]++

		// Check if the current element is the majority element
		if frequencyMap[num] > len(nums)/2 {
			return num
		}
	}

	// The majority element always exists, so this line should not be reached.
	return -1
}

func MajorityElement(nums []int) int {
	majority := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			count++
		} else {
			count--
			if count == 0 {
				majority = nums[i]
				count = 1
			}
		}
	}
	return majority
}
