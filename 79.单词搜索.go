/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
// var direction = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
// var ret bool

func exist(board [][]byte, word string) bool {
	// ret = false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if backtracking(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func backtracking(board [][]byte, word string, i, j, index int) bool {
	if i < 0 || i >= len(board) {
		return false
	}
	if j < 0 || (i < len(board) && j >= len(board[i])) {
		return false
	}
	if board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 && board[i][j] == word[index] {
		return true
	}
	board[i][j] += 100
	if backtracking(board, word, i+1, j, index+1) ||
		backtracking(board, word, i-1, j, index+1) ||
		backtracking(board, word, i, j+1, index+1) ||
		backtracking(board, word, i, j-1, index+1) {
		return true
	}
	board[i][j] -= 100
	return false
}

// @lc code=end

