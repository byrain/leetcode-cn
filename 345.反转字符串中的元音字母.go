/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func isVowels(l string) bool {
	if l == "a" || l == "e" || l == "i" || l == "o" || l == "u" || l == "A" || l == "E" || l == "I" || l == "O" || l == "U" {
		return true
	}
	return false
}

func reverseVowels(s string) string {
	sList := []rune{}
	for _, v := range s {
		sList = append(sList, v)
	}
	p1 := 0
	p2 := len(sList) - 1
	for {
		if p1 >= p2 {
			break
		}
		if isVowels(string(sList[p1])) && isVowels(string(sList[p2])) {
			sList[p1], sList[p2] = sList[p2], sList[p1]
			p1++
			p2--
			continue
		}
		if !isVowels(string(sList[p1])) && isVowels(string(sList[p2])) {
			p1++
			continue
		}
		if isVowels(string(sList[p1])) && !isVowels(string(sList[p2])) {
			p2--
			continue
		}
		if !isVowels(string(sList[p1])) && !isVowels(string(sList[p2])) {
			p1++
			p2--
		}
	}

	r := ""
	for _, v := range sList {
		r += string(v)
	}
	return r
}

// @lc code=end

