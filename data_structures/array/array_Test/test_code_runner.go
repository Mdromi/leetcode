package arrayTest

import (
	"fmt"
	"testing"

	Array "github.com/Mdromi/dsa-cp/data_structures/array"
)

func RunInsertElementTest(t *testing.T, numbers []int, indexingValue int, targetedIndex int, afterIndexingValue []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("Targeted Index: %d\n", targetedIndex)

	// Call the function to insert the value at the specified index
	result := Array.InsertElementOfArrayAnyPosition(numbers, indexingValue, targetedIndex)

	fmt.Printf("\nResult: %v\n", result)

	// Check if the result is equal to the expected result
	if fmt.Sprintf("%v", result) == fmt.Sprintf("%v", afterIndexingValue) {
		fmt.Printf("\x1b[32m%s\x1b[0m\n", "✔ Test Passed: Value inserted at the correct index")
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", "✘ Test Failed: Target value not inserted at the correct index")
		t.Errorf("Expected array %v, but got %v", afterIndexingValue, result)
	}
}
