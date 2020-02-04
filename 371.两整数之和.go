/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	for {
		if b != 0 {
			temp := a ^ b
			b = (a & b) << 1
			a = temp
		} else {
			break
		}
	}
	return a
}

// @lc code=end

