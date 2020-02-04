import "math"

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	min := math.MaxInt32
	for _, v := range nums {
		if v <= min {
			min = v
		}
	}
	r := 0
	for _, v := range nums {
		r += (v - min)
	}
	return r
}

// @lc code=end

