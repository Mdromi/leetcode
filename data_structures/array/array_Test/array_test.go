package arrayTest

import (
	"fmt"
	"testing"

	Array "github.com/Mdromi/dsa-cp/data_structures/array"
)

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
