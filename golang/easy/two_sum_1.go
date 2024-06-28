package main

import "fmt"

// Function to find two indices such that nums[i] + nums[j] = target
func twoSum(nums []int, target int) []int {
    // Create a map to store numbers and their indices
    numMap := make(map[int]int)

    // Iterate through the array
    for i, num := range nums {
        complement := target - num
        // Check if the complement exists in the map
        if index, found := numMap[complement]; found {
            return []int{index, i}
        }
        // If not found, add the current number and its index to the map
        numMap[num] = i
    }

    // Solution is guaranteed to exist, so we should never reach here
    return nil
}

func main() {
    // Example 1
    nums1 := []int{2, 7, 11, 15}
    target1 := 9
    fmt.Println(twoSum(nums1, target1)) // [0, 1]

    // Example 2
    nums2 := []int{3, 2, 4}
    target2 := 6
    fmt.Println(twoSum(nums2, target2)) // [1, 2]

    // Example 3
    nums3 := []int{3, 3}
    target3 := 6
    fmt.Println(twoSum(nums3, target3)) // [0, 1]
}
