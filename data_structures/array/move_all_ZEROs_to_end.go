package array

func MoveAllZEROSToEnd(array []int) []int {
	n := len(array)
	count := 0

	for i := 0; i < n; i++ {
		if array[i] != 0 {
			// Swap non-zero element with the first zero encountered
			array[i], array[count] = array[count], array[i]
			count++
		}
	}

	return array
}
