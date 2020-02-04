/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	z := 1
	t := num
	count := 0
	for ; t > 0; t = t / 2 {
		count++
	}
	for {
		if count-1 >= 0 {
			z *= 2
			count--
		} else {
			break
		}
	}
	return num ^ (z - 1)
}

// @lc code=end

