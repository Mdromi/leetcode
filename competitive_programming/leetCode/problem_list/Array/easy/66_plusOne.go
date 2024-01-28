package easy

/**
Problem: https://leetcode.com/problems/plus-one/
*/

/*
Time: O(n)
Space: O(1)
*/
func PlusOne(digits []int) []int {
	n := len(digits)
	carry := 1

	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
func PlusOne2(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
/*
Time: O(n)
Space: O(1)
*/
func PlusOne_MySolution(digits []int) []int {
	len := len(digits) - 1

	if digits[len] != 9 {
		digits[len] = digits[len] + 1
		return digits
	}

	for i := len; i >= 0; i-- {
		if i == 0 && digits[i] == 9 {
			digits[i] = 1
			digits = append(digits, 0)
		} else if digits[i-1] != 9 {
			digits[i] = 0
			digits[i-1] += 1
			break
		} else {
			digits[i] = 0
		}
	}

	return digits
}
