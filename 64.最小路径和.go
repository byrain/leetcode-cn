/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				grid[0][j] = grid[0][j-1] + grid[0][j]
			}
			if i != 0 && j == 0 {
				grid[i][0] = grid[i-1][0] + grid[i][0]
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

// @lc code=end

