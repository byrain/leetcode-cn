/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		vv := target - v
		for i, j := range nums[k+1:] {
			if vv == j {
				return []int{k, i + k + 1}
			}
		}
	}
	return []int{}
}

// @lc code=end

// hash
// func twoSum(nums []int, target int) []int {
// 	t := map[int][]int{}
// 	for k, v := range nums {
// 		if _, ok := t[v]; !ok {
// 			t[v] = []int{k}
// 		} else {
// 			t[v] = append(t[v], k)
// 		}
// 	}

// 	for k, v := range t {
// 		tarNum := target - k
// 		if _, ok := t[tarNum]; ok {
// 			if tarNum == k && len(v) > 1 {
// 				return v
// 			}
// 			if tarNum != k {
// 				return []int{v[0], t[tarNum][0]}
// 			}
// 		}
// 	}
// 	return []int{}
// }