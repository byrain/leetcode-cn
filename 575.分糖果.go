/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candies []int) int {
	m := map[int]struct{}{}
	count := 0
	for _, v := range candies {
		if _, ok := m[v]; !ok {
			count++
			m[v] = struct{}{}
		}
	}
	r := count
	if count > len(candies)/2 {
		r = len(candies) / 2
	}
	return r
}

// @lc code=end

