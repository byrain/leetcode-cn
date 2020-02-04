/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p := len(nums1) - 1
	p1 := m - 1
	p2 := n - 1
	for {
		if p < 0 || p2 < 0 {
			return
		}
		if p1 < 0 || nums1[p1] <= nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
			p--
		} else {
			nums1[p] = nums1[p1]
			p1--
			p--
		}
	}
}

// @lc code=end

