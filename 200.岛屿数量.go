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
	if i >= len(grid) {
		return
	}
	if i < len(grid) && j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	used[i][j] = true
	// up
	if i-1 >= 0 && grid[i-1][j] == '1' && !used[i-1][j] {
		dfs(i-1, j, grid)
	}
	// down
	if i+1 < len(grid) && grid[i+1][j] == '1' && !used[i+1][j] {
		dfs(i+1, j, grid)
	}
	// left
	if j-1 >= 0 && grid[i][j-1] == '1' && !used[i][j-1] {
		dfs(i, j-1, grid)
	}
	// right
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' && !used[i][j+1] {
		dfs(i, j+1, grid)
	}
}

// @lc code=end

