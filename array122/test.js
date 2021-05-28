/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  if (prices.length < 2) {
    return 0;
  }
  let dp = new Array(prices.length).fill([0, 0]);
  dp[0] = [0, -prices[0]];

  for (let i = 1; i < prices.length; i++) {
    dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
    dp[i][1] = Math.max(dp[i - 1][0] - prices[i], dp[i - 1][1]);
  }

  return dp[prices.length - 1][0];
};

console.log(maxProfit([7, 1, 5, 3, 6, 4]));
