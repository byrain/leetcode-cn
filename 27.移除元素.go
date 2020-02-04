/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	pSlow := 0
	pFast := 0
	for {
		if pFast == len(nums) {
			break
		}
		if nums[pFast] != val {
			nums[pSlow], nums[pFast] = nums[pFast], nums[pSlow]
			pSlow += 1

		}
		pFast += 1
	}
	return pSlow
}

// @lc code=end

