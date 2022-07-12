/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

var used map[int][]bool

// @lc code=start
func numIslands(grid [][]byte) int {
	count := 0
	used = map[int][]bool{}
	for i := 0; i < len(grid); i++ {
		used[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			used[i][j] = false
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !used[i][j] && grid[i][j] == '1' {
				count++
				dfs(i, j, grid)
			}
		}
	}
	return count
}

func dfs(i, j int, grid [][]byte) {
	if i < 0 || i >= len(grid) {
		return
	}
	if j < 0 || (i < len(grid) && j >= len(grid[i])) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	if used[i][j] {
		return
	}
	used[i][j] = true
	dfs(i-1, j, grid)
	dfs(i+1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)
}

// @lc code=end

