/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func partition(nums []int, left, right int) int {
	key := nums[right]
	i := left - 1

	for j := left; j < right; j++ {
		if nums[j] <= key {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[right] = nums[right], nums[i+1]

	return i + 1
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

// @lc code=end

