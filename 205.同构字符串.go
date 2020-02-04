/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if s == t {
		return true
	}
	mapping1 := map[byte]byte{}
	mapping2 := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := mapping1[s[i]]; ok {
			if mapping1[s[i]] == t[i] {
				continue
			}
			return false
		}
		if _, ok := mapping2[t[i]]; ok {
			if mapping2[t[i]] == s[i] {
				continue
			}
			return false
		}
		mapping1[s[i]] = t[i]
		mapping2[t[i]] = s[i]
	}
	return true
}

// @lc code=end

