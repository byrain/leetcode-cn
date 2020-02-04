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

	lenMin := math.MaxInt64
	for _, v := range strs {
		if lenMin > len(v) {
			lenMin = len(v)
		}
	}

	commonStr := ""

	for i := 0; i < lenMin; i++ {
		for _, v := range strs {
			if v[i] != commonStr {
				
			}
		}
		commonStr = 
	}
	return ""
}

// @lc code=end

