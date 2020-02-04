/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	p1 := 0
	p2 := len(s) - 1
	for {
		if p1 >= p2 {
			break
		}
		s[p1], s[p2] = s[p2], s[p1]
		p1++
		p2--
	}
}

// @lc code=end

