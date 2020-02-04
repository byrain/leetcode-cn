/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pFast := 1
	pSlow := 0
	for {
		if pFast >= len(nums) {
			break
		}
		if nums[pFast] > nums[pSlow] {
			pSlow += 1
			nums[pSlow] = nums[pFast]
		}
		pFast += 1
	}
	return pSlow + 1
}

// @lc code=end

