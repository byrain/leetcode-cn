/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	dict := map[int]bool{}
	r := []int{}
	for _, v := range nums1 {
		dict[v] = true
	}
	for _, v := range nums2 {
		if _, ok := dict[v]; ok {
			if dict[v] {
				r = append(r, v)
				dict[v] = false
			}
		}
	}
	return r
}

// @lc code=end

