import "math/bits"

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	return bits.OnesCount(uint(z))
}

// @lc code=end

