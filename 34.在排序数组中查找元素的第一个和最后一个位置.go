/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left := 0
	leftP := -1
	right := len(nums) - 1
	for left <= right {
		p := left + (right-left)/2
		if nums[p] == target {
			leftP = p
			right = p - 1
		} else if target < nums[p] {
			right = p - 1
		} else {
			left = p + 1
		}
	}

	rightP := -1
	right = len(nums) - 1
	for left <= right {
		p := left + (right-left)/2
		if nums[p] == target {
			rightP = p
			left = p + 1
		} else if target < nums[p] {
			right = p - 1
		} else {
			left = p + 1
		}
	}

	return []int{leftP, rightP}
}

// @lc code=end

