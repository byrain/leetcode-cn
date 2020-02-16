/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := map[int]int{}
	dp[1] = 1
	dp[2] = 2
	i := 3
	for i < n {
		dp[i] = dp[i-1] + dp[i-2]
		i++
	}
	return dp[n-1] + dp[n-2]
}

// @lc code=end

