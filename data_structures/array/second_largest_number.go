package array

func SecondLargestNumber(array []int) int {
	if len(array) < 2 {
		return 0
	}

	largest := array[0]
	secondLargest := array[0]

	for _, num := range array {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest && num < largest {
			secondLargest = num
		}
	}

	return secondLargest
}
