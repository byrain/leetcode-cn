import "strings"

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	r := 0
	ss := strings.Split(s, " ")
	for _, v := range ss {
		if v != "" {
			r += 1
		}
	}
	return r
}

// @lc code=end

