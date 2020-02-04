/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xReal := x
	xRev := 0
	for {
		if x == 0 {
			break
		}
		xRev = xRev*10 + x%10
		x = x / 10
	}

	if xRev == xReal {
		return true
	}
	return false
}

// @lc code=end

