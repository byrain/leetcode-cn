/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	max := 0
	i, j := 0, 1
	for j < len(nums) {
		if nums[j-1] >= nums[j] {
			tempMax := j - i
			if tempMax > max {
				max = tempMax
			}
			i = j
			j++
		} else if nums[j-1] < nums[j] {
			j++
		}
	}
	tempMax := j - i
	if tempMax > max {
		max = tempMax
	}
	return max
}

// @lc code=end

