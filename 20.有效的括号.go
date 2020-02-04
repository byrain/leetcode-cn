import "strings"

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	for {
		if strings.Contains(s, "()") ||
			strings.Contains(s, "[]") ||
			strings.Contains(s, "{}") {
			s = strings.Replace(s, "()", "", -1)
			s = strings.Replace(s, "[]", "", -1)
			s = strings.Replace(s, "{}", "", -1)
		} else {
			break
		}
	}

	return s == ""
}

// @lc code=end

