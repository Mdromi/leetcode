package main

import "fmt"

func majorityElement(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    // Variables to store the potential candidates and their counts
    var candidate1, candidate2, count1, count2 int
    
    // First pass: Find potential candidates
    for _, num := range nums {
        if num == candidate1 {
            count1++
        } else if num == candidate2 {
            count2++
        } else if count1 == 0 {
            candidate1 = num
            count1 = 1
        } else if count2 == 0 {
            candidate2 = num
            count2 = 1
        } else {
            count1--
            count2--
        }
    }
    
    // Reset counts for validation
    count1, count2 = 0, 0
    
    // Second pass: Validate the candidates
    for _, num := range nums {
        if num == candidate1 {
            count1++
        } else if num == candidate2 {
            count2++
        }
    }
    
    // Collect the valid candidates
    result := []int{}
    if count1 > len(nums)/3 {
        result = append(result, candidate1)
    }
    if count2 > len(nums)/3 {
        result = append(result, candidate2)
    }
    
    return result
}

func main() {
    // Test cases
    nums1 := []int{3, 2, 3}
    fmt.Println(majorityElement(nums1)) // Output: [3]

    nums2 := []int{1}
    fmt.Println(majorityElement(nums2)) // Output: [1]

    nums3 := []int{1, 2}
    fmt.Println(majorityElement(nums3)) // Output: [1, 2]
}
