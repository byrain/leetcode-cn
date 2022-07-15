/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	subBoard := map[int]map[byte]bool{}
	for i := 0; i < 9; i++ {
		subBoard[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		row := make(map[byte]bool, 9)
		column := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if row[board[i][j]] == true && board[i][j] != '.' {
				return false
			}
			row[board[i][j]] = true

			if column[board[j][i]] == true && board[j][i] != '.' {
				return false
			}
			column[board[j][i]] = true

			if subBoard[i/3*3+j/3][board[i][j]] == true && board[i][j] != '.' {
				return false
			}
			subBoard[i/3*3+j/3][board[i][j]] = true
		}
	}
	return true
}

// @lc code=end

