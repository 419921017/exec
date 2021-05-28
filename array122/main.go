package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(myMaxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	fmt.Printf("%v\n", dp)

	for i := 1; i < len(prices); i++ {
		// 第i天交易完后手里没有股票的最大利润
		// dp[i-1][0], 前一天已经没有股票
		// dp[i-1][1], 前一天手里有股票, 如果要卖出, 获得 prices[i] 的收益
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		// 第i天交易完后手里持有一支的最大利润
		// dp[i-1][0], 前一天没股票, 需要购买股票, 扣除 prices[i] 收益
		// dp[i-1][1], 前一天已经有一只股票
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])

	}
	fmt.Printf("%v\n", dp)

	// 最后持有股票的收益一定小于不持有股票的收益
	// dp[n-1][0] > dp[n-1][1]
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func myMaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 没股票
		// dp[i-1][0]没股票
		// dp[i-1][1]+prices[i], 有股票卖了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 有股票
		// dp[i-1][0]-prices[i], 有股票卖了
		// dp[i-1][1], 有股票没卖
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
