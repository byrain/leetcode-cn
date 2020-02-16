/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	// i, j := 0, 1
	// count := 1
	// ret := 0
	// for i < len(s) {
	// 	if s[i] == s[j] {
	// 		j++
	// 		count++
	// 	} else {
	// 		if count > 0 {
	// 			count--
	// 			j++
	// 		}
	// 	}
	// }
	// return ret
	prev := 0
	cur := 1
	sum := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			if prev < cur {
				sum += prev
			} else {
				sum += cur
			}
			prev, cur = cur, 1
		}
	}
	if prev < cur {
		sum += prev
	} else {
		sum += cur
	}
	return sum
}

// @lc code=end

