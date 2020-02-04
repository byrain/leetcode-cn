/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	r := []int{}
	for i := 1; i < len(nums)+1; i++ {
		r = append(r, i)
	}
	for i := 0; i < len(nums); i++ {
		if r[nums[i]-1] == nums[i] {
			r[nums[i]-1] = 0
		}
	}
	nums = []int{}
	for i := 0; i < len(r); i++ {
		if r[i] != 0 {
			nums = append(nums, r[i])
		}
	}

	return nums
}

// @lc code=end

