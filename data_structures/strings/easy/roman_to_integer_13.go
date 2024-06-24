package main

import "fmt"

// Function to convert Roman numeral to integer
func romanToInt(s string) int {
    // Map to store Roman numerals and their values
    romanMap := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    // Variable to store the result
    total := 0
    n := len(s)

    // Iterate through the string
    for i := 0; i < n; i++ {
        // Get the value of the current Roman numeral
        current := romanMap[s[i]]

        // Check the value of the next Roman numeral (if any)
        if i+1 < n && current < romanMap[s[i+1]] {
            // If the current value is less than the next value, subtract it
            total -= current
        } else {
            // Otherwise, add the current value
            total += current
        }
    }

    return total
}

func main() {
    // Example 1
    s1 := "III"
    fmt.Println(romanToInt(s1)) // 3

    // Example 2
    s2 := "LVIII"
    fmt.Println(romanToInt(s2)) // 58

    // Example 3
    s3 := "MCMXCIV"
    fmt.Println(romanToInt(s3)) // 1994
}
