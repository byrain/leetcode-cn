/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	m := map[string]struct{}{}
	for _, v := range J {
		m[string(v)] = struct{}{}
	}

	ret := 0
	for _, letter := range S {
		if _, ok := m[string(letter)]; ok {
			ret++
		}
	}
	return ret
}

// @lc code=end

