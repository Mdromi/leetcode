package array

/**
QUESTION: You are given an array of integers. An element in the array is considered a "leader" if it is greater than all the elements to its right. Your task is to implement a function that finds and return all the leaders in the given array.
*/

func LeadersInAnArray(array []int) []int {
	n := len(array)

	if n == 0 {
		return []int{} // Return an empty array for an empty input array
	}

	leaders := make([]int, 0)
	maxRight := array[n-1]

	// The rightmost element is always a leader
	leaders = append(leaders, maxRight)

	for i := n - 2; i >= 0; i-- {
		if array[i] > maxRight {
			maxRight = array[i]
			leaders = append(leaders, maxRight)
		}
	}

	// Reverse the leaders array to maintain the original order
	reverse(leaders)

	return leaders
}

func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}
