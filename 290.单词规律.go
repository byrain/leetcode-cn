import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	if len(strList) != len(pattern) {
		return false
	}
	d1 := map[byte]string{}
	d2 := map[string]byte{}
	for i := 0; i < len(strList); i++ {
		if _, ok := d1[pattern[i]]; !ok {
			d1[pattern[i]] = strList[i]
		} else {
			if d1[pattern[i]] != strList[i] {
				return false
			}
		}

		if _, ok := d2[strList[i]]; !ok {
			d2[strList[i]] = pattern[i]
		} else {
			if d2[strList[i]] != pattern[i] {
				return false
			}
		}
	}
	return true
}

// @lc code=end

