/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ret := []string{}
	if n <= 0 {
		return ret
	}
	dfs("", n, n, &ret)
	return ret
}

func dfs(curString string, left, right int, ret *[]string) {
	if left == 0 && right == 0 {
		*ret = append(*ret, curString)
		return
	}

	if left > right {
		return
	}

	if left > 0 {
		dfs(curString+"(", left-1, right, ret)
	}

	if right > 0 {
		dfs(curString+")", left, right-1, ret)
	}
}

// @lc code=end

