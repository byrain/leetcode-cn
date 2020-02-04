/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n < 1 {
		return false
	}
	for {
		if n%3 != 0 {
			return false
		}
		n = n / 3
		if n == 1 {
			return true
		}
	}
	return false
}

// @lc code=end

