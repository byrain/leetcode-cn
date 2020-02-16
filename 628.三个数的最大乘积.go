/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	maxA, maxB, maxC := -1000, -1000, -1000
	minA, minB := 1000, 1000
	for _, v := range nums {
		if v > maxA {
			maxC = maxB
			maxB = maxA
			maxA = v
		} else if v > maxB {
			maxC = maxB
			maxB = v
		} else if v > maxC {
			maxC = v
		}
		if v < minA {
			minB = minA
			minA = v
		} else if v < minB {
			minB = v
		}
	}
	max1 := maxA * maxB * maxC
	max2 := minA * minB * maxA
	if max1 > max2 {
		return max1
	}
	return max2
}

// @lc code=end

