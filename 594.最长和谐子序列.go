/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	res := 0
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	for k, v := range count {
		if count[k-1] != 0 && res < count[k-1]+v {
			res = count[k-1] + v
		}
		if count[k+1] != 0 && res < count[k+1]+v {
			res = count[k+1] + v
		}
	}
	return res
}

// @lc code=end

