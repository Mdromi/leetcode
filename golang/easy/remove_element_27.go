package main

import "fmt"

func removeElement(nums []int, val int) int {
    k := 0 // Initialize the pointer `k`

    for _, num := range nums {
        if num != val {
            nums[k] = num
            k++
        }
    }

    return k
}

func main() {
    // Test cases
    nums1 := []int{3, 2, 2, 3}
    val1 := 3
    k1 := removeElement(nums1, val1)
    fmt.Println(k1, nums1[:k1]) // Output: 2, [2 2]

    nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
    val2 := 2
    k2 := removeElement(nums2, val2)
    fmt.Println(k2, nums2[:k2]) // Output: 5, [0 1 3 0 4]
}
