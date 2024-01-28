package easy

/**
Problem: https://leetcode.com/problems/single-number/
*/

/*
Time: O(n)
Space: O(n)
*/

func SingleNumber_usingHashMap(nums []int) int {
	frequencyMap := make(map[int]int)

	// Count the frequency of each element
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Find the element with frequency 1
	for key, value := range frequencyMap {
		if value == 1 {
			return key
		}
	}

	return -1 // Return -1 if no single element is found (shouldn't happen based on the problem statement)
}

/*
Time: O(n)
Space: O(1)
*/
func SingleNumber(nums []int) int {
	result := 0

	// XOR all elements in the array
	for _, num := range nums {
		result ^= num
	}

	return result
}
