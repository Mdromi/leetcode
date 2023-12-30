package array

func LargestNumber(array []int) int {
	if len(array) == 0 {
		// Return an appropriate value, like 0 or an error, depending on your use case
		return 0
	}

	largest := array[0]

	for _, num := range array {
		if num > largest {
			largest = num
		}
	}

	return largest
}
