/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	commonStr := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		curStr := strs[i]
		for {
			if j == len(curStr) || j == len(commonStr) {
				if len(curStr) < len(commonStr) {
					commonStr = curStr
				}
				break
			}
			if curStr[j] != commonStr[j] {
				commonStr = commonStr[:j]
				break
			}
			j++
		}
	}
	return commonStr
}

// @lc code=end

