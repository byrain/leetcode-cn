/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	dp := map[int]int{0: cost[0], 1: cost[1]}
	for i := 2; i < len(cost); i++ {
		temp := dp[i-1]
		if dp[i-1] > dp[i-2] {
			temp = dp[i-2]
		}
		dp[i] = temp + cost[i]
	}
	if dp[len(cost)-1] < dp[len(cost)-2] {
		return dp[len(cost)-1]
	}
	return dp[len(cost)-2]
}

// @lc code=end

