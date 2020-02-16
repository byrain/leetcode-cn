import "math"

/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(S string, C byte) []int {
	ret := []int{}
	nums := []int{}
	for k, v := range S {
		if byte(v) == C {
			nums = append(nums, k)
		}
	}
	for k, _ := range S {
		dist := 10000
		for _, v := range nums {
			t := int(math.Abs(float64(v - k)))
			if t < dist {
				dist = t
			}
		}
		ret = append(ret, dist)
	}
	return ret
}

// @lc code=end

