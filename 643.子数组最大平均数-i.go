import "math"

/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt32
	if k == len(nums) {
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
		return float64(maxSum) / float64(k)
	}
	for i := 0; i <= len(nums)-k; i++ {
		sum := nums[i]
		for j := 1; j < k; j++ {
			sum = sum + nums[i+j]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}

// @lc code=end

