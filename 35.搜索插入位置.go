/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			return k
		}
	}
	return len(nums)
}

// @lc code=end

