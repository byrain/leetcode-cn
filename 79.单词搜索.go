/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
var direction = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var ret bool

func exist(board [][]byte, word string) bool {
	ret = false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				backtracking(board, word, i, j, 0)
			}
		}
	}
	return ret
}

func backtracking(board [][]byte, word string, i, j, index int) {
	if i < 0 || i >= len(board) {
		return
	}
	if j < 0 || (i < len(board) && j >= len(board[i])) {
		return
	}

	if board[i][j] != word[index] {
		return
	}
	if index == len(word)-1 && board[i][j] == word[index] {
		ret = true
		return
	}
	for _, d := range direction {
		board[i][j] += 100
		backtracking(board, word, i+d[0], j+d[1], index+1)
		board[i][j] -= 100
	}
	return
}

// @lc code=end

