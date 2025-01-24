package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
func MaxProfit(prices []int) int {

	lessPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		profit := price - lessPrice
		if profit > maxProfit {
			maxProfit = profit
		}

		if price < lessPrice {
			lessPrice = price
		}
	}

	return maxProfit
}
