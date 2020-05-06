import "fmt"

/*
 * @lc app=leetcode.cn id=999 lang=golang
 *
 * [999] 车的可用捕获量
 */

// @lc code=start
func numRookCaptures(board [][]byte) int {
	i := 0
	j := 0
	for _, row := range board {
		for _, col := range board {
			if board[row][col] == 'R' {
				i = row
				j = col
				break
			}
		}
	}
	fmt.Println(i, j)
	return 0
}

// @lc code=end

