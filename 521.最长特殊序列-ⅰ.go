/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] 最长特殊序列 Ⅰ
 */

// @lc code=start
func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	}
	if len(a) < len(b) {
		return len(b)
	}

	if a == b {
		return -1
	}
	return len(a)

}

// @lc code=end

