/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	t := n ^ (n >> 1)
	return t&(t+1) == 0
}

// @lc code=end

