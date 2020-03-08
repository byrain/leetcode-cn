/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	j := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// @lc code=end

