import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	if s == "" || s == " " {
		return true
	}
	runeList := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			runeList = append(runeList, v)
		}
	}
	r := ""
	for i := 0; i < len(runeList); i++ {
		r = r + string(runeList[i])
	}

	r1 := ""
	for j := len(runeList) - 1; j >= 0; j-- {
		r1 = r1 + string(runeList[j])
	}

	return strings.ToUpper(r) == strings.ToUpper(r1)
}

// @lc code=end

