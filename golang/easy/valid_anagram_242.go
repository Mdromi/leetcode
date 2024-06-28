package main

import "fmt"

// Function to check if t is an anagram of s
func isAnagram(s string, t string) bool {
    // If lengths are not the same, they can't be anagrams
    if len(s) != len(t) {
        return false
    }

    // Create a frequency map to count occurrences of each character in s
    charCount := make(map[rune]int)

    // Increment the count for each character in s
    for _, char := range s {
        charCount[char]++
    }

    // Decrement the count for each character in t
    for _, char := range t {
        charCount[char]--
        // If the count goes below zero, it means t has an extra character
        if charCount[char] < 0 {
            return false
        }
    }

    // If we go through all characters without issue, they are anagrams
    return true
}

func main() {
    // Example 1
    s1 := "anagram"
    t1 := "nagaram"
    fmt.Println(isAnagram(s1, t1)) // true

    // Example 2
    s2 := "rat"
    t2 := "car"
    fmt.Println(isAnagram(s2, t2)) // false
}
