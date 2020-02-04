/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for {
		t := num
		if num%2 == 0 {
			num = num / 2
		}
		if num%3 == 0 {
			num = num / 3
		}
		if num%5 == 0 {
			num = num / 5
		}

		if num <= 5 {
			return true
		}

		if t == num {
			return false
		}
	}
}

// @lc code=end

