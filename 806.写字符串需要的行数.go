/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, S string) []int {
	line := 1
	width := 0
	for _, v := range S {
		p := v - 'a'
		w := widths[p]
		width += w
		if width > 100 {
			line++
			width = w
		}
	}
	return []int{line, width}
}

// @lc code=end

