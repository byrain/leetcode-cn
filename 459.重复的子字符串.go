import "strings"

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	return strings.Index(s[1:]+s[:len(s)-1], s) >= 0
}

// @lc code=end

