/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	s1 := map[rune]int{}
	for _, v := range s {
		s1[v] += 1
	}
	t1 := map[rune]int{}
	for _, v := range t {
		t1[v] += 1
	}

	if len(s1) != len(t1) {
		return false
	}
	for k, v := range s1 {
		if t1[k] != v {
			return false
		}
	}

	return true
}

// @lc code=end

