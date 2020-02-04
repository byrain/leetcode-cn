import "sort"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)
	return (float64((nums1[l/2] + nums1[(l-1)/2])) / 2)
}

// @lc code=end

