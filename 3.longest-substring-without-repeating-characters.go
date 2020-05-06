/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	str := ""
	strMap := make(map[rune]bool, len(s))
	max := 0
	for _, v := range s {
		if ok := strMap[v]; ok {
			for k, vTemp := range str {
				strMap[vTemp] = false
				if vTemp == v {
					str = str[k+1:]
					strMap[vTemp] = true
					break
				}
			}
		}
		str = str + string(v)
		if max < len(str) {
			max = len(str)
		}
		strMap[v] = true
	}

	return max
}

// @lc code=end

