import "sort"

/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	dupNum := 0
	sort.Ints(nums)
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i-1]^nums[i] == 0 {
			dupNum = nums[i]
		}
	}
	sumAll := (1 + len(nums)) * len(nums) / 2
	return []int{dupNum, sumAll - sum + dupNum}
}

// @lc code=end

