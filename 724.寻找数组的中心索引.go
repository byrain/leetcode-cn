/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	nums = append(nums, 0)
	nums = append([]int{0}, nums...)
	leftSum := 0
	for i := 0; i < len(nums)-2; i++ {
		leftSum += nums[i]
		if leftSum*2+nums[i+1] == sum {
			return i
		}
	}
	return -1
}

// @lc code=end

