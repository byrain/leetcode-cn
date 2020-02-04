/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	if k >= len(nums) {
		k = k - len(nums)
	}
	nums = reverse(nums, 0, len(nums)-1)
	nums = reverse(nums, 0, k-1)
	nums = reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, s, e int) []int {
	for {
		if s >= e {
			return nums
		}
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}

// @lc code=end

