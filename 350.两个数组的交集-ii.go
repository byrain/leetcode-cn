/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	dict1 := map[int]int{}
	dict2 := map[int]int{}
	r := []int{}
	for _, v := range nums1 {
		dict1[v] += 1
	}
	for _, v := range nums2 {
		dict2[v] += 1
	}
	for k, v := range dict1 {
		if _, ok := dict2[k]; ok {
			m := 0
			if dict2[k] < v {
				m = dict2[k]
			} else {
				m = v
			}
			for i := 0; i < m; i++ {
				r = append(r, k)
			}
			// dict2[k] = 0
		}
	}
	return r
}

// @lc code=end

