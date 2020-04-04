import "sort"

/*
 * @lc app=leetcode.cn id=945 lang=golang
 *
 * [945] 使数组唯一的最小增量
 */

// @lc code=start
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	moves := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			pre := A[i]
			A[i] = A[i-1] + 1
			moves += A[i] - pre
		}
	}
	return moves
}

// @lc code=end

