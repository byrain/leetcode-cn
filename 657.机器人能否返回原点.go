/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	m := map[rune]int{
		'U': 0,
		'D': 0,
		'L': 0,
		'R': 0,
	}
	for _, v := range moves {
		m[v] += 1
	}

	return m['U'] == m['D'] && m['R'] == m['L']
}

// @lc code=end

