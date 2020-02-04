/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	count := 0
	for {
		if n > 0 {
			n = n / 5
			count += n
		} else {
			return count
		}
	}

	return count
}

// @lc code=end

