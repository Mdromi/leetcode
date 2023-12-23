package arrayTest

import (
	"fmt"
	"reflect"
	"testing"
)

// Common function to print the test result
func PrintTestResult(t *testing.T, functionName string, result, expected []int) {
	fmt.Printf("\nResult of %s: %v\n", functionName, result)
	fmt.Printf("Expected: %v\n", expected)

	if reflect.DeepEqual(result, expected) {
		fmt.Printf("\x1b[32m%s\x1b[0m\n", "✔ Test Passed: Array operation completed correctly")
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", "✘ Test Failed: Array operation did not produce the expected result")
		t.Errorf("Expected array %v, but got %v", expected, result)
	}
}
