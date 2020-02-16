/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	i := 0
	for {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
		i++
	}
	return 0
}

// @lc code=end

