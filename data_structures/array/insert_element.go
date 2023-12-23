package array

import "fmt"

func InsertElementOfArrayAnyPosition(array []int, value int, IndexingPosition int) []int {
	// Ensure the position is within the array bounds
	if IndexingPosition < 0 || IndexingPosition > len(array) {
		fmt.Println("Invalid index position")
		return array
	}

	// Use the custom append function
	array = customAppend(array, 0)

	// Shift elements to make space for the new value
	for i := len(array) - 1; i > IndexingPosition; i-- {
		array[i] = array[i-1]
	}

	// Insert the new value at the specified position
	array[IndexingPosition] = value

	return array
}
func customAppend(arr []int, value int) []int {
	// Check if there's enough capacity
	if len(arr) == cap(arr) {
		// If not, double the capacity
		newCapacity := cap(arr) * 2
		newArr := make([]int, len(arr), newCapacity)
		copy(newArr, arr)
		arr = newArr
	}

	// Append the new value
	arr = arr[:len(arr)+1]
	arr[len(arr)-1] = value

	return arr
}
