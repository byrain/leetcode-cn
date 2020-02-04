import "sort"

/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	begin := -1
	end := -1
	i, j := 0, len(nums)-1
	for {
		if begin >= 0 && end >= 0 {
			break
		}
		if i > j {
			break
		}
		if nums[i] != temp[i] {
			begin = i
		}
		if begin < 0 {
			i++
		}

		if nums[j] != temp[j] {
			end = j
		}
		if end < 0 {
			j--
		}
	}
	if begin < 0 || end < 0 {
		return 0
	}
	return end - begin + 1

}

// @lc code=end

