package main

import "fmt"

func isPalindrome(x int) bool {
    // Special cases:
    // 1. Negative numbers are not palindromes
    // 2. Numbers ending in 0 (except 0 itself) are not palindromes
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    
    // To store the reversed second half of x
    reversed := 0
    
    // Reverse the second half of x
    for x > reversed {
        reversed = reversed * 10 + x % 10
        x /= 10
    }
    
    // x is a palindrome if reversed equals to x (for even length) or reversed/10 equals to x (for odd length)
    return x == reversed || x == reversed/10
}

func main() {
    // Test cases
    fmt.Println(isPalindrome(121))   // Output: true
    fmt.Println(isPalindrome(-121))  // Output: false
    fmt.Println(isPalindrome(10))    // Output: false
}
