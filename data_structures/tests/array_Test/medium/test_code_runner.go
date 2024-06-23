package arrayTest

import (
	"fmt"
	"reflect"
	"testing"

	Medium "github.com/Mdromi/leetcode/data_structures/array/medium"
)

// 15. 3Sum
func RunFourSum(t *testing.T, array []int, target int, expected [][]int) {
	fmt.Printf("\nOriginal Array: %v\n", array)

	result := Medium.FourSum(array, target)

	successMessage := "Three Four Sum is correct."
	failureMessage := "Three Four Sum is incorrect."

	if reflect.DeepEqual(result, expected) {
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "✔ Test Passed:", successMessage)
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m %s\n", "✘ Test Failed:", failureMessage)
		t.Errorf("Expected: %v, but got %v", expected, result)
	}
}

// 15. 3Sum
func RunThreeSum(t *testing.T, array []int, expected [][]int) {
	fmt.Printf("\nOriginal Array: %v\n", array)

	result := Medium.ThreeSum(array)

	successMessage := "Three Sum is correct."
	failureMessage := "Three Sum is incorrect."

	if reflect.DeepEqual(result, expected) {
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "✔ Test Passed:", successMessage)
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m %s\n", "✘ Test Failed:", failureMessage)
		t.Errorf("Expected: %v, but got %v", expected, result)
	}
}

// 11. Container With Most Water
func RunMaxArea(t *testing.T, height []int, maxArea int) {
	fmt.Printf("\nOriginal Array: %v\n", height)

	result := Medium.MaxArea(height)

	successMessage := "MaxArea  is correct."
	failureMessage := "MaxArea  number is incorrect."

	if result == maxArea {
		fmt.Printf("\x1b[32m%s\x1b[0m %s\n", "✔ Test Passed:", successMessage)
	} else {
		fmt.Printf("\x1b[31m%s\x1b[0m %s\n", "✘ Test Failed:", failureMessage)
		t.Errorf("Expected: %v, but got %v", maxArea, result)
	}
}
