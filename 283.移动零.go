/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	p1 := 0
	p2 := 0
	for {
		if p2 >= len(nums) {
			break
		}
		if nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1 += 1
			p2 += 1
		} else {
			p2++
		}
	}
}

// @lc code=end

