/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := 1
	candidate := nums[0]
	for i := 1; i < len(nums); i++ {
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
		if count <= 0 {
			candidate = nums[i]
			count = 1
		}

	}
	return candidate
}

// @lc code=end

