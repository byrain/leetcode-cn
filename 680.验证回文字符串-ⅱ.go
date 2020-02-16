/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {
	s1 := ""
	s2 := ""
	i, j := 0, len(s)-1
	ret := true
	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			s1 = s[i+1 : j+1]
			s2 = s[i:j]
			break
		}
	}
	if !isValidPalindrome(s1) && !isValidPalindrome(s2) {
		ret = false
	}
	return ret
}

func isValidPalindrome(s string) bool {
	i, j := 0, len(s)-1
	ret := true
	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			ret = false
			break
		}
	}
	return ret
}

// @lc code=end

