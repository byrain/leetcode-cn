import "strings"

/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	if word == strings.ToUpper(word) {
		return true
	}
	if word == strings.ToLower(word) {
		return true
	}
	if string(word[0]) == strings.ToUpper(string(word[0])) {
		if word[1:] == strings.ToLower(word[1:]) {
			return true
		}
	}
	return false
}

// @lc code=end

