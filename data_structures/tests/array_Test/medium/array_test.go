package arrayTest

import (
	"testing"
)

// go test -run

func TestFourSum(t *testing.T) {
	test1 := []int{1, 0, -1, 0, -2, 2}
	test2 := []int{2, 2, 2, 2, 2}

	result1 := [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
	result2 := [][]int{{2, 2, 2, 2}}

	RunFourSum(t, test1, 0, result1)
	RunFourSum(t, test2, 8, result2)
}

func TestThreeSum(t *testing.T) {
	test1 := []int{-1, 0, 1, 2, -1, -4}
	test2 := []int{0, 1, 1}
	test3 := []int{0, 0, 0}

	result1 := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result2 := [][]int{} // Empty slice of slices
	result3 := [][]int{{0, 0, 0}}

	RunThreeSum(t, test1, result1)
	RunThreeSum(t, test2, result2)
	RunThreeSum(t, test3, result3)
}

func TestMaxProfit(t *testing.T) {
	// Test Case 1: getting the max area
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := 49

	// Ren the test
	RunMaxArea(t, height, maxArea)
}
