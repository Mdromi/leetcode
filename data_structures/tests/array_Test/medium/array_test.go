package arrayTest

import (
	"testing"
)

// go test -run

func TestMaxProfit(t *testing.T) {
	// Test Case 1: getting the max area
	height := []int{1,8,6,2,5,4,8,3,7}
	maxArea := 49

	// Ren the test
	RunMaxArea(t, height, maxArea)
}

