package array

func SearchElementIndex(searchList []int, targetValue int) int {
	for i, v := range searchList {
		if v == targetValue {
			return i
		}
	}
	return -1
}
