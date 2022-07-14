/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
var ret []string

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return nil
	}
	ret = []string{}
	dfs([]byte(S), 0)
	return ret
}

func dfs(S []byte, index int) {
	if len(S) == index {
		ret = append(ret, string(S))
		return
	}
	dfs(S, index+1)
	if !(S[index] < 'A' || S[index] > 'z') {
		S[index] ^= 32
		dfs(S, index+1)
	}
}

// @lc code=end

