package array

func LeftRotateArrayByOne(arr []int) []int {
	// Check if the array is empty or has only one element
	if len(arr) <= 1 {
		return arr
	}

	// Store the first element in a temporary variable
	temp := arr[0]

	// Shift all elements to the left
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	// Place the temporary variable at the end
	arr[len(arr)-1] = temp

	return arr
}
