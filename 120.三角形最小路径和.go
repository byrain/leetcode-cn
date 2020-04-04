import "math"

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, len(triangle))
	for k, v := range triangle {
		dp[k] = make([]int, len(v))
	}
	triangle[0][0] = triangle[0][0]
	triangle[1][0] = triangle[1][0] + triangle[0][0]
	triangle[1][1] = triangle[1][1] + triangle[0][0]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][0] = triangle[i-1][0] + triangle[i][0]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	ret := math.MaxInt32
	for _, v := range triangle[len(triangle)-1] {
		ret = min(ret, v)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

