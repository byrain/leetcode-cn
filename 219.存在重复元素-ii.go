/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	i := 0
	ret := false
	window := []int{}
	for {
		if i >= len(nums) {
			break
		}
		if i+k+1 > len(nums) {
			window = nums[i+1:]
		} else {
			window = nums[i+1 : i+k+1]
		}

		for _, v := range window {
			if nums[i] == v {
				return true
			}
		}
		i++
	}
	return ret
}

// @lc code=end

