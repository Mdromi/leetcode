import (
    "fmt"
    "sort"
)
// 1. Sorting the Array
func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}

// 2. Using a Hash Map
func containsDuplicate2(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}


func main() {
    fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // true
    fmt.Println(containsDuplicate2([]int{1, 2, 3, 4})) // false
}
