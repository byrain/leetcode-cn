/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	temp := 0
	for _, v := range nums {
		if v == 1 {
			temp++
			if temp > max {
				max = temp
			}
		} else {
			temp = 0
		}
	}
	return max
}

// @lc code=end

