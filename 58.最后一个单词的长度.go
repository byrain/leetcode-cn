import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	l := len(s)
	if l <= 0 {
		return 0
	}
	kTemp := 0
	for k, v := range s {
		if string(v) == " " {
			if k == (l - 1) {
				l = l - 1
				break
			}
			kTemp = k + 1
		}
	}
	return l - kTemp
}

// @lc code=end

