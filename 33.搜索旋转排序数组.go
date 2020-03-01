/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		p := left + (right-left)/2
		if target == nums[p] {
			return p
		}

		if p == len(nums)-1 {
			break
		}

		if nums[p+1] <= nums[right] {
			if target >= nums[p+1] && target <= nums[right] {
				left = p + 1
			} else {
				right = p - 1
			}
		} else {
			if target >= nums[left] && target <= nums[p] {
				right = p - 1
			} else {
				left = p + 1
			}
		}
	}
	return -1
}

// @lc code=end

