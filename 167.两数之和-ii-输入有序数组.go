/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		if i >= j {
			return []int{}
		}
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

		if numbers[j] > target && numbers[j] != 0 {
			j--
			continue
		}
		if numbers[i]+numbers[j] > target {
			j--
			continue
		}

		if numbers[i]+numbers[j] < target {
			i++
			continue
		}

		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

// @lc code=end

