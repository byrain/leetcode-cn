/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	i := len(nums) - 2
	j := i + 1
	for i >= 0 {
		if nums[i] >= nums[j] {
			i--
			j--
		} else {
			break
		}
	}

	for k := len(nums) - 1; k >= j; k-- {
		if i < 0 {
			break
		}
		if nums[k] > nums[i] {
			t := nums[i]
			nums[i] = nums[k]
			nums[k] = t
			break
		}
	}

	m := j
	n := len(nums) - 1
	for m <= n {
		t := nums[m]
		nums[m] = nums[n]
		nums[n] = t
		m++
		n--
	}
}

// @lc code=end

