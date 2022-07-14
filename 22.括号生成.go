/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
var ret []string

func generateParenthesis(n int) []string {
	ret := []string{}
	backtracking(n, []byte{}, 0, 0)
	return ret
}

func backtracking(n int, path []byte, leftBracket, rightBracket int) {
	if len(path) == 2*n && leftBracket == rightBracket {
		newPath := make([]byte, len(path))
		copy(newPath, path)
		ret = append(ret, string(newPath))
		return
	}
	if leftBracket < rightBracket {
		return
	}
	if leftBracket < n {
		path = append(path, '(')
		backtracking(n, path, leftBracket+1, rightBracket)
		path = path[:len(path)-1]
	}
	if rightBracket < n {
		path = append(path, ')')
		backtracking(n, path, leftBracket, rightBracket+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end

