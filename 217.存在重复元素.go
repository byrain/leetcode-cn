/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = struct{}{}
		}
	}
	return false
}

// @lc code=end

