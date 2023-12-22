package array

import "fmt"

func DeleteElementOfArrayAnyPosition(array []int, deletedPosition int) []int {
	// Ensure the position is within the array bounds
	if deletedPosition < 0 || deletedPosition >= len(array) {
		fmt.Println("Invalid index position")
		return array
	}

	// Shift elements to remove the element at the specified position
	for index := deletedPosition; index < len(array)-1; index++ {
		array[index] = array[index+1]
	}

	// Resize the array by slicing to exclude the last element
	array = array[:len(array)-1]

	return array
}
