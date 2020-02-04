import "strings"

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	r := ""
	sList := strings.Split(s, " ")
	for _, v := range sList {
		sList := []byte(v)
		for i := 0; i < len(sList)/2; i++ {
			sList[i], sList[len(sList)-i-1] = sList[len(sList)-i-1], sList[i]
		}
		r += string(sList) + " "
	}
	return strings.TrimSuffix(r, " ")
}

// @lc code=end

