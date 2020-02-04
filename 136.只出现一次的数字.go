/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	t := 0
	for _, v := range nums {
		t = v ^ t
	}
	return t
}

// @lc code=end

