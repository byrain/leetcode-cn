/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x
	ret := 0
	for {
		p := (left + right) / 2
		if (p+1)*(p+1) > x && p*p <= x {
			ret = p
			break
		}
		if p*p > x {
			right = p - 1
		} else {
			left = p + 1
		}
	}
	return ret
}

// @lc code=end

