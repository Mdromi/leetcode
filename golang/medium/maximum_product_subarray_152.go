package main

import "fmt"

func maxProduct(nums []int) int {
    // Initialize maxProduct, currentMax, and currentMin to the first element
    maxProduct := nums[0]
    currentMax := nums[0]
    currentMin := nums[0]
    
    // Iterate through the array starting from the second element
    for i := 1; i < len(nums); i++ {
        num := nums[i]
        
        // If the current number is negative, swap currentMax and currentMin
        if num < 0 {
            currentMax, currentMin = currentMin, currentMax
        }
        
        // Calculate the new currentMax and currentMin
        currentMax = max(num, currentMax * num)
        currentMin = min(num, currentMin * num)
        
        // Update maxProduct if currentMax is greater
        maxProduct = max(maxProduct, currentMax)
    }
    
    return maxProduct
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    // Test cases
    fmt.Println(maxProduct([]int{2, 3, -2, 4})) // Output: 6
    fmt.Println(maxProduct([]int{-2, 0, -1}))  // Output: 0
}
