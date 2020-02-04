/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}
	if num < 1 {
		return false
	}
	for {
		if num%4 != 0 {
			return false
		}
		num = num / 4
		if num == 1 {
			return true
		}
	}
	return false
}

// @lc code=end

