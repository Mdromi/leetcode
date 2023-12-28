package array

import "sort"

func RemoveDuplicatesValuesFromArray(array []int) []int {
	// Create a map to store unique values
	uniqueValues := make(map[int]bool)

	// Create a new array to store the result without duplicates
	result := []int{}

	// Iterate through the input array
	for _, value := range array {
		// Check if the value is not present in the uniqueValues map
		if !uniqueValues[value] {
			// Add the value to the result array
			result = append(result, value)

			// Mark the value as seen in the map
			uniqueValues[value] = true
		}
	}

	return result
}

func RemoveDuplicatesValuesFromArray_saveSpaceComplexity(array []int) []int {
	// Check if the array is empty or has only one element
	if len(array) < 2 {
		return array
	}

	// Sort the array to group duplicate values together
	sort.Ints(array)

	// Use two pointers to iterate through the array
	writeIndex := 1

	for i := 1; i < len(array); i++ {
		// If the current element is different from the previous one,
		// update the array and move the write index
		if array[i] != array[i-1] {
			array[writeIndex] = array[i]
			writeIndex++
		}
	}

	// Truncate the array to remove the duplicates at the end
	return array[:writeIndex]
}
