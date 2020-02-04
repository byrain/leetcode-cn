/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 {
		return false
	}
	for {
		t := n % 2
		if t == 1 {
			return false
		}
		n = n / 2
		if n == 1 {
			return true
		}
	}
	return false
}

// @lc code=end

