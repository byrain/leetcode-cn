/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
<<<<<<< Updated upstream
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

=======
	res := []string{}
	if n == 0 {
		return res
	}

	dfs("", n, n, &res)
	return res
}

func dfs(curStr string, left, right int, res *[]string) {
	// 因为每一次尝试，都使用新的字符串变量，所以无需回溯
	// 在递归终止的时候，直接把它添加到结果集即可，注意与「力扣」第 46 题、第 39 题区分
	if left == 0 && right == 0 {
		*res = append(*res, curStr)
		return
	}

	// 剪枝（如图，左括号可以使用的个数严格大于右括号可以使用的个数，才剪枝，注意这个细节）
>>>>>>> Stashed changes
	if left > right {
		return
	}

	if left > 0 {
<<<<<<< Updated upstream
		dfs(curString+"(", left-1, right, ret)
	}

	if right > 0 {
		dfs(curString+")", left, right-1, ret)
=======
		dfs(curStr+"(", left-1, right, res)
	}

	if right > 0 {
		dfs(curStr+")", left, right-1, res)
>>>>>>> Stashed changes
	}
}

// @lc code=end

