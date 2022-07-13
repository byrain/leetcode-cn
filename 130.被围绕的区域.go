/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
	}
	for i := 1; i < len(board[0])-1; i++ {
		dfs(board, 0, i)
		dfs(board, len(board)-1, i)
	}
	for i := 1; i < len(board[0])-1; i++ {
		dfs(board, 0, i)
		dfs(board, len(board)-1, i)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) {
		return
	}
	if j < 0 || (i < len(board) && j >= len(board[i])) {
		return
	}
	if board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
	return
}

// @lc code=end

