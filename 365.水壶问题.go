/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 */

// @lc code=start
func canMeasureWater(x int, y int, z int) bool {
	if x == z || y == z {
		return true
	}

	if x+y < z {
		return false
	}

	if z%(gcd(x, y)) == 0 {
		return true
	}
	return false
}

func gcd(x, y int) int {
	temp := x % y
	if temp > 0 {
		return gcd(y, temp)
	} else {
		return y
	}
}

// @lc code=end

