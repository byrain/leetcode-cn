/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums2) == 0 {
		return []int{}
	}
	r := []int{}
	m := map[int]int{}
	i := 0
	j := 1
	for {
		if i == len(nums2)-1 {
			m[nums2[i]] = -1
			break
		}
		if nums2[j] > nums2[i] {
			m[nums2[i]] = nums2[j]
			i++
			j = i + 1
		} else {
			j++
			if j >= len(nums2) {
				m[nums2[i]] = -1
				i++
				j = i + 1
			}
		}

	}
	for _, v := range nums1 {
		r = append(r, m[v])
	}
	return r
}

// @lc code=end

