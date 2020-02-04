/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	n := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] = digits[i] + 1
			if digits[i] > 9 {
				digits[i] = 0
				n = 1
			}
			if n == 0 {
				return digits
			}
			continue
		}

		if n > 0 {
			digits[i] = digits[i] + n
			if digits[i] > 9 {
				digits[i] = 0
				n = 1
			} else {
				n = 0
			}

			continue
		}

		if n == 0 {
			return digits
		}

		if i == 0 {
			return append([]int{1}, digits...)
		}
	}
	if n > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end

