package array

func MaximumDifference(array []int) int {
	n := len(array)

	if n < 2 {
		return 0 // Return 0 for insufficient elements
	}

	maxDiff := array[1] - array[0]
	minElement := array[0]

	for i := 1; i < n; i++ {
		if array[i] < minElement {
			minElement = array[i]
		} else if array[i]-minElement > maxDiff {
			maxDiff = array[i] - minElement
		}
	}

	return maxDiff
}
