/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	max := 0
	second := 0
	r := -1

	for k, v := range nums {
		if v > max {
			second = max
			max = v
			r = k
		} else if v > second {
			second = v
		}
	}
	if max >= second*2 {
		return r
	}
	return -1
}

// @lc code=end

