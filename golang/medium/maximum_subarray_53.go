package main

import "fmt"

func maxSubArray(nums []int) int {
    // Initialize currentSum and maxSum to the first element of the array
    currentSum := nums[0]
    maxSum := nums[0]
    
    // Iterate through the array starting from the second element
    for i := 1; i < len(nums); i++ {
        // Update currentSum to be the maximum of the current element
        // or the sum of currentSum and the current element
        currentSum = max(nums[i], currentSum + nums[i])
        // Update maxSum if currentSum is greater
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // Test cases
    fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // Output: 6
    fmt.Println(maxSubArray([]int{1}))                            // Output: 1
    fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))               // Output: 23
}
