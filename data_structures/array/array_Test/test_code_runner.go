package arrayTest

import (
	"fmt"
	"reflect"
	"testing"

	Array "github.com/Mdromi/dsa-cp/data_structures/array"
)

func RunSecondLargestNumber(t *testing.T, numbers []int, secondLargestNumber int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("The second largest number is: %d\n", secondLargestNumber)

	// Test LargestNumber
	result := Array.SecondLargestNumber(numbers)
	fmt.Printf("The SecondLargestNumber Function Return This Number: %d\n", result)

	successMessage := "Second Largest number is correct."
	failureMessage := "Second Largest number is incorrect."

	if result == secondLargestNumber {
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "✔ Test Passed:", successMessage)
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m %s\n", "✘ Test Failed:", failureMessage)
		t.Errorf("Expected: %v, but got %v", secondLargestNumber, result)
	}
}

func RunLargestNumber(t *testing.T, numbers []int, largestNumber int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("The largest number is: %d\n", largestNumber)

	// Test LargestNumber
	result := Array.LargestNumber(numbers)
	fmt.Printf("The LargestNumber Function Return This Number: %d\n", result)

	successMessage := "Largest number is correct."
	failureMessage := "Largest number is incorrect."

	if result == largestNumber {
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "✔ Test Passed:", successMessage)
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m %s\n", "✘ Test Failed:", failureMessage)
		t.Errorf("Expected: %v, but got %v", largestNumber, result)
	}
}

// RunRemoveDuplicatesValuesFromArray runs tests for removing duplicates value
func RunRemoveDuplicatesValuesFromArray(t *testing.T, numbers []int, expected []int) {
	// Remove duplicates from the array
	result1 := Array.RemoveDuplicatesValuesFromArray(numbers)
	expected2 := []int{1, 7, 9, 2, 3, 4, 5, 6, 8, 10}

	// Test the removal of duplicates
	PrintTestResult(t, "RemoveDuplicatesValuesFromArray", result1, expected2, "Duplicates removed successfully", "Duplicates were not removed as expected")

	// Apply left rotation to the array (modify the function name if needed)
	result2 := Array.RemoveDuplicatesValuesFromArray_saveSpaceComplexity(result1)

	// Test the left rotation
	PrintTestResult(t, "RemoveDuplicatesValuesFromArray_saveSpaceComplexity", result2, expected, "Duplicates removed with space optimization", "Duplicates were not removed with space optimization as expected")
}

// RunLeftRotateArrayByOneTest runs tests for left rotating an array by one position
func RunLeftRotateArrayByOneTest(t *testing.T, numbers []int, expected []int) {
	// Make a copy of the original array to preserve it
	originalNumbers := make([]int, len(numbers))
	copy(originalNumbers, numbers)

	// Left rotate the array by 1 position
	result := Array.LeftRotateArrayByOne(numbers)

	// Print the test result
	PrintTestResult(t, "LeftRotateArrayByOne", result, expected, "Array left rotated by one position", "Array did not left rotate as expected")

	// Ensure the original array is preserved
	if !reflect.DeepEqual(numbers, originalNumbers) {
		fmt.Println("\x1b[31mWarning:\x1b[0m Original array is not preserved.")
	}
}

// RunReverseElementTest runs tests for array reversal functions
func RunReverseElementTest(t *testing.T, numbers []int, afterReversingArray []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)

	// Test ReverseTheArraySaveSpaceComplexity
	resultSpaceComplexity := Array.ReverseTheArraySaveSpaceComplexity(numbers)
	PrintTestResult(t, "ReverseTheArraySaveSpaceComplexity", resultSpaceComplexity, afterReversingArray, "Array reversed successfully", "Array did not reverse as expected")

	// Test ReverseTheArrayXOR
	resultXOR := Array.ReverseTheArrayXOR(numbers)
	PrintTestResult(t, "ReverseTheArrayXOR", resultXOR, afterReversingArray, "Array reversed successfully", "Array did not reverse as expected")

	// Test ReverseTheArrayLinearTime
	resultLinearTime := Array.ReverseTheArrayLinearTime(numbers)
	PrintTestResult(t, "ReverseTheArrayLinearTime", resultLinearTime, afterReversingArray, "Array reversed successfully", "Array did not reverse as expected")
}

// RunDeleteElementTest runs tests for deleting an element from the array
func RunDeleteElementTest(t *testing.T, numbers []int, deletedIndex int, afterDeletingValue []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("Deleted Index: %d\n", deletedIndex)

	// Call the function to delete the value at the specified index
	result := Array.DeleteElementOfArrayAnyPosition(numbers, deletedIndex)

	fmt.Printf("\nResult: %v\n", result)

	// Print the test result
	PrintTestResult(t, "DeleteElementOfArrayAnyPosition", result, afterDeletingValue, "Value deleted at the correct index", "Target value not deleted at the correct index")
}

// RunInsertElementTest runs tests for inserting an element into the array
func RunInsertElementTest(t *testing.T, numbers []int, indexingValue int, targetedIndex int, afterIndexingValue []int) {
	fmt.Printf("\nOriginal Array: %v\n", numbers)
	fmt.Printf("Targeted Index: %d\n", targetedIndex)

	// Call the function to insert the value at the specified index
	result := Array.InsertElementOfArrayAnyPosition(numbers, indexingValue, targetedIndex)

	fmt.Printf("\nResult: %v\n", result)

	// Print the test result
	PrintTestResult(t, "InsertElementOfArrayAnyPosition", result, afterIndexingValue, "Value inserted at the correct index", "Target value not inserted at the correct index")
}
