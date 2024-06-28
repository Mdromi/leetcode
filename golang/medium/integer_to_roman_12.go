package main

import "fmt"

func intToRoman(num int) string {
    // Arrays of Roman symbols and their corresponding values
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    
    // Variable to store the result
    result := ""

    // Iterate through the values array
    for i := 0; i < len(values); i++ {
        // While num is greater than or equal to the current value
        while num >= values[i] {
            result += symbols[i] // Append the symbol to the result
            num -= values[i]     // Subtract the value from num
        }
    }

    return result
}

func intToRoman_2(num int) string {
    ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
    tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    hrns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    ths := []string{"", "M", "MM", "MMM"}

    return ths[num/1000] + hrns[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}

func main() {
    // Example 1
    num1 := 3749
    fmt.Println(intToRoman(num1)) // MMMDCCXLIX

    // Example 2
    num2 := 58
    fmt.Println(intToRoman(num2)) // LVIII

    // Example 3
    num3 := 1994
    fmt.Println(intToRoman(num3)) // MCMXCIV
}
