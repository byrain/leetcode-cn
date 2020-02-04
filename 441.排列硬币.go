/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	r := 0
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum == n {
			r = i
			break
		}
		if sum > n {
			r = i - 1
			break
		}
	}
	return r
}

// @lc code=end

