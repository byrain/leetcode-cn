import "math"

/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		v := i*i + j*j
		if v == c {
			return true
		} else if v < c {
			i++
		} else {
			j--
		}
	}
	return false
}

// @lc code=end

