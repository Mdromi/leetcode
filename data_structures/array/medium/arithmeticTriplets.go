package medium

func ArithmeticTriplets(nums []int, diff int) int {
	count := 0
	vis := make(map[int]bool)

	for _, x := range nums {
		vis[x] = true
	}

	for _, x := range nums {
		if vis[x+diff] && vis[x+diff*2] {
			count++
		}
	}

	return count
}
