/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	marked := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		marked[i] = make([]bool, cols)
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			count++
			dfs(i, j, rows, cols, grid, marked)
		}
	}
}

func dfs(i, j, rows, cols int, grid [][]byte, marked [][]bool) {

}

// @lc code=end

