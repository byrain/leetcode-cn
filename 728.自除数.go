/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	ret := []int{}

	for i := left; i <= right; i++ {
		temp := i
		isSelfDividing := true
		for temp > 0 {
			p := temp % 10
			if p == 0 || i%p != 0 {
				isSelfDividing = false
				break
			}
			temp = temp / 10
		}
		if isSelfDividing {
			ret = append(ret, i)
		}
	}
	return ret
}

// @lc code=end

