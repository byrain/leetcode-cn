import "sort"

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] > 0 {
		return 0
	}
	l := len(nums)
	if nums[l-1] != l {
		return nums[l-1] + 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return 0
}

// @lc code=end

