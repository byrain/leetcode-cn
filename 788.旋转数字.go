/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 */

// @lc code=start
func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		n := i
		flag := false
		for n > 0 {
			x := n % 10
			if x == 3 || x == 4 || x == 7 {
				break
			}
			if flag || x == 2 || x == 5 || x == 6 || x == 9 {
				flag = true
			}
			n /= 10
		}
		if n == 0 && flag {
			count++
		}
	}
	return count
}

// @lc code=end

