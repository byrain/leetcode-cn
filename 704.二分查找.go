/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		pivot := (left + right) / 2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}

	}
	return -1
}

// @lc code=end

