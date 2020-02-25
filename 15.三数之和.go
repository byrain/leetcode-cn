import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// 剔除重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			// 剔除重复值
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			// 剔除重复值
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				temp := []int{nums[i], nums[j], nums[k]}
				ret = append(ret, temp)
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return ret
}

// @lc code=end

