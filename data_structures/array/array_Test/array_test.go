package arrayTest

import (
	"fmt"
	"testing"

	Array "github.com/Mdromi/dsa-cp/data_structures/array"
)

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
