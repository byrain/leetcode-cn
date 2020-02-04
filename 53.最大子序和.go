/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	sum := 0
	maxSubSum := nums[0]
	for _, v := range nums {
		sum = sum + v
		if sum > maxSubSum {
			maxSubSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSubSum
}

// @lc code=end

