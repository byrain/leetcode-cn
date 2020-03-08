/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if target == matrix[0][0] || target == matrix[m-1][n-1] {
		return true
	}
	for i := 0; i < m; i++ {
		if target == matrix[i][n-1] || target == matrix[i][0] {
			return true
		}
		if target < matrix[i][n-1] && target > matrix[i][0] {
			for j := 1; j < n; j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}

// @lc code=end

