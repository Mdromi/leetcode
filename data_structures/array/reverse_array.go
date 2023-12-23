package array

// Reverse the array in-place with O(1) space complexity
func ReverseTheArraySaveSpaceComplexity(array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

// Reverse the array in-place with O(1) space complexity and potential time complexity improvement
func ReverseTheArrayXOR(array []int) []int {
	// Check if the array is empty or has only one element
	if len(array) <= 1 {
		return array
	}

	// Reverse the array using XOR swapping without using extra space
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		// XOR swap without using additional variable
		array[i] ^= array[j]
		array[j] ^= array[i]
		array[i] ^= array[j]
	}

	return array
}

func ReverseTheArrayLinearTime(array []int) []int {
	// Check if the array is empty or has only one element
	if len(array) <= 1 {
		return array
	}

	// Initialize two pointers at the beginning and end of the array
	i, j := 0, len(array)-1

	// Swap elements using a temporary variable until the pointers meet
	for i < j {
		// Swap array[i] and array[j]
		temp := array[i]
		array[i] = array[j]
		array[j] = temp

		// Move the pointers towards the center
		i++
		j--
	}

	return array
}
