package arrayTest

import (
	"fmt"
	"reflect"
	"testing"
)

// Common function to print the test result
func PrintTestResult(t *testing.T, functionName string, result, expected []int, successMessage, failureMessage string) {
	if reflect.DeepEqual(result, expected) {
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "✔ Test Passed:", successMessage)
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m %s\n", "✘ Test Failed:", failureMessage)
		t.Errorf("Function: %s | Expected: %v, but got %v", functionName, expected, result)
	}
}
