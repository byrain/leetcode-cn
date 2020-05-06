/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 */

// @lc code=start

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	ret := false

	if min(rec1[2], rec2[2]) > max(rec1[0], rec2[0]) &&
		min(rec1[3], rec2[3]) > max(rec1[1], rec2[1]) {
		ret = true
	}

	return ret
}

// @lc code=end

