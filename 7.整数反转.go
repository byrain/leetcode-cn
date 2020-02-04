import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	r := 0
	for {
		if x == 0 {
			break
		}
		r = r*10 + x%10
		x = x / 10
	}

	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return r
}

// @lc code=end

