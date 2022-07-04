/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 时间复杂度 O(n2)
	// for k, v := range nums {
	// 	vv := target - v
	// 	for i, j := range nums[k+1:] {
	// 		if vv == j {
	// 			return []int{k, i + k + 1}
	// 		}
	// 	}
	// }
	// return []int{}

	// 时间复杂度 O(n)
	m := map[int][]int{}
	for k, v := range nums {
		m[v] = append(m[v], k)
	}
	for k, v := range m {
		t := target - k
		if _, ok := m[t]; ok {
			if t != k {
				return []int{v[0], m[t][0]}
			}
			if len(m[k]) > 1 {
				return m[k][0:2]
			}
		}
	}
	return []int{}
}

// @lc code=end
