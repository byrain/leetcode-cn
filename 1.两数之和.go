/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		vv := target - v
		for i, j := range nums[k+1:] {
			if vv == j {
				return []int{k, i + k + 1}
			}
		}
	}
	return []int{}
}

// @lc code=end

