package array

/**
Question:
You are given an array representing the stock prices on different days. Write a function in C++ to find the maximum profit that can be obtained by buying and selling the stock, considering you can make multiple transactions. The function signature is:

MaxProfit( price []int,  start, end int) int
{1, 5, 3, 8, 12}, start = 0, end = 4;
Maximum Profit in Stock Market = 13

The function should return the maximum profit that can be achieved within the given range [start, end] of the stock prices. The input array price contains integers representing the stock prices on different days.


Note:
In the given example, the maximum profit is obtained by buying the stock on day 1 (price = 1) and selling it on day 2 (price = 5), then buying again on day 3 (price = 3) and selling on day 4 (price = 8), and finally buying on day 5 (price = 12) and selling. The total profit is (5-1) + (8-3) + (12-3) = 13.

*/

func MaxProfit(price []int) int {
	maxProfit := 0

	for i := 1; i < len(price); i++ {
		// If the current price is higher than the previous one, add the profit
		if price[i] > price[i-1] {
			maxProfit += price[i] - price[i-1]
		}
	}

	return maxProfit
}
