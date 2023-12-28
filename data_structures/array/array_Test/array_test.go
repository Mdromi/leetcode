package arrayTest

import (
	"fmt"
	"testing"

	Array "github.com/Mdromi/dsa-cp/data_structures/array"
)

// go test -run

func TestRemoveDuplicatesValuesFromArray(t *testing.T) {
	// Test Case 1: Left rotate the array by 1 position
	numbers1 := []int{1, 7, 7, 9, 9, 2, 3, 4, 5, 6, 7, 8, 9, 10, 8, 9, 10, 1, 5, 4, 1}

	// Define the expected array
	afterRemovingTheDuplicatesValue := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Ren the test
	RunRemoveDuplicatesValuesFromArray(t, numbers1, afterRemovingTheDuplicatesValue)
}

func TestLeftRotateArrayByOne(t *testing.T) {
	// Test Case 1: Left rotate the array by 1 position
	numbers1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Define the expected result after left rotating the array by 1 position
	afterLeftRotatingArray := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	// Run the test
	RunLeftRotateArrayByOneTest(t, numbers1, afterLeftRotatingArray)
}

func TestReverseTheArraySaveSpaceComplexity(t *testing.T) {
	// Test Case 1: Delete at the start
	numbers1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

	afterRreversingArray := []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	RunReverseElementTest(t, numbers1, afterRreversingArray)
}

func TestDeleteElementOfArrayAnyPosition(t *testing.T) {
	// Test Case 1: Delete at the start
	numbers1 := []int{5, 20, 28, 290, 30}
	deletedIndex1 := 0
	afterDeletingValue1 := []int{20, 28, 290, 30}
	RunDeleteElementTest(t, numbers1, deletedIndex1, afterDeletingValue1)

	// Test Case 2: Delete in the middle
	numbers2 := []int{5, 20, 28, 290, 30}
	deletedIndex2 := 2
	afterDeletingValue2 := []int{5, 20, 290, 30}
	RunDeleteElementTest(t, numbers2, deletedIndex2, afterDeletingValue2)

	// Test Case 3: Delete at the end
	numbers3 := []int{5, 20, 28, 290, 30}
	deletedIndex3 := len(numbers3) - 1
	afterDeletingValue3 := []int{5, 20, 28, 290}
	RunDeleteElementTest(t, numbers3, deletedIndex3, afterDeletingValue3)

	// Test Case 4: Delete insertion index
	numbers4 := []int{5, 20, 28, 290, 30}
	deletedIndex4 := 3
	afterDeletingValue4 := []int{5, 20, 28, 30}
	RunDeleteElementTest(t, numbers4, deletedIndex4, afterDeletingValue4)
}

func TestInsertElementOfArrayAnyPosition(t *testing.T) {
	// Test Case 1: Insert at the start
	numbers1 := []int{5, 20, 28, 290, 30}
	indexingValue1 := 45
	targetedIndex1 := 0
	afterIndexingValue1 := []int{45, 5, 20, 28, 290, 30}
	RunInsertElementTest(t, numbers1, indexingValue1, targetedIndex1, afterIndexingValue1)

	// Test Case 2: Insert in the middle
	numbers2 := []int{5, 20, 28, 290, 30}
	indexingValue2 := 45
	targetedIndex2 := 2
	afterIndexingValue2 := []int{5, 20, 45, 28, 290, 30}
	RunInsertElementTest(t, numbers2, indexingValue2, targetedIndex2, afterIndexingValue2)

	// Test Case 3: Insert at the end
	numbers3 := []int{5, 20, 28, 290, 30}
	indexingValue3 := 45
	targetedIndex3 := len(numbers3)
	afterIndexingValue3 := []int{5, 20, 28, 290, 30, 45}
	RunInsertElementTest(t, numbers3, indexingValue3, targetedIndex3, afterIndexingValue3)

	// Test Case 4: Random insertion index
	numbers4 := []int{5, 20, 28, 290, 30}
	indexingValue4 := 45
	targetedIndex4 := 3
	afterIndexingValue4 := []int{5, 20, 28, 45, 290, 30}
	RunInsertElementTest(t, numbers4, indexingValue4, targetedIndex4, afterIndexingValue4)
}

func TestSearchElement(t *testing.T) {
	numbers := []int{5, 20, 28, 290, 30}
	targetValue := 28
	expectedIndex := 2

	result := Array.SearchElementIndex(numbers, targetValue)

	if result == expectedIndex {
		fmt.Printf("Value %d found at index %d\n", targetValue, result)
	} else {
		fmt.Println("Target value not found")
		t.Errorf("Expected index %d, but got %d", expectedIndex, result)
	}
}
