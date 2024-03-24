package arrayTest

import (
	"fmt"
	"testing"

	Medium "github.com/Mdromi/leetcode/data_structures/array/medium"
)

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
