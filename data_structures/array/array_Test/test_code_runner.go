package arrayTest

import (
	"fmt"
	"testing"

	Array "github.com/Mdromi/dsa-cp/data_structures/array"
)

// RunReverseElementTest runs tests for array reversal functions
func RunReverseElementTest(t *testing.T, numbers []int, afterReversingArray []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)

	// Test ReverseTheArraySaveSpaceComplexity
	resultSpaceComplexity := Array.ReverseTheArraySaveSpaceComplexity(numbers)
	PrintTestResult(t, "ReverseTheArraySaveSpaceComplexity", resultSpaceComplexity, afterReversingArray)

	// Test ReverseTheArrayXOR
	resultXOR := Array.ReverseTheArrayXOR(numbers)
	fmt.Println("numbers", numbers)
	PrintTestResult(t, "ReverseTheArrayXOR", resultXOR, numbers)

	// Test ReverseTheArrayLinearTime
	resultLinearTime := Array.ReverseTheArrayLinearTime(numbers)
	PrintTestResult(t, "ReverseTheArrayLinearTime", resultLinearTime, afterReversingArray)
}

// RunDeleteElementTest runs tests for deleting an element from the array
func RunDeleteElementTest(t *testing.T, numbers []int, deletedIndex int, afterDeletingValue []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("Deleted Index: %d\n", deletedIndex)

	// Call the function to delete the value at the specified index
	result := Array.DeleteElementOfArrayAnyPosition(numbers, deletedIndex)

	fmt.Printf("\nResult: %v\n", result)

	// Print the test result
	PrintTestResult(t, "DeleteElementOfArrayAnyPosition", result, afterDeletingValue)
}

// RunInsertElementTest runs tests for inserting an element into the array
func RunInsertElementTest(t *testing.T, numbers []int, indexingValue int, targetedIndex int, afterIndexingValue []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("Targeted Index: %d\n", targetedIndex)

	// Call the function to insert the value at the specified index
	result := Array.InsertElementOfArrayAnyPosition(numbers, indexingValue, targetedIndex)

	fmt.Printf("\nResult: %v\n", result)

	// Print the test result
	PrintTestResult(t, "InsertElementOfArrayAnyPosition", result, afterIndexingValue)
}
