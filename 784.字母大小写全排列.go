/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
func letterCasePermutation(S string) []string {
	ret := []string{S}
	S = S + "0"
	for k, letter := range S {
		if letter >= 'a' && letter <= 'z' {
			ret = append(ret, S[:k]+string()+S[k+1:])
		}
		if letter >= 'A' && letter <= 'Z' {

		}
	}
	return ret
}

// @lc code=end

