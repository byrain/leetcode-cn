/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	window := map[byte]int{}
	need := map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	needCount := 0

	min := len(s) + 1
	left, right := 0, 0
	ret := [2]int{0, 0}
	for right < len(s) {
		window[s[right]]++
		if window[s[right]] == need[s[right]] {
			needCount++
		}
		for needCount == len(need) {
			if right-left < min {
				min = right - left
				ret[0] = left
				ret[1] = right + 1
			}
			if window[s[left]] == need[s[left]] {
				needCount--
			}
			window[s[left]]--
			left++
		}
		right++
	}
	if needCount != len(need) && left == 0 {
		return ""
	}
	return s[ret[0]:ret[1]]
}

// @lc code=end

