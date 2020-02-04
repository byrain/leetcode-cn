import "math"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	r := 0
	minList := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	isExist := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, ok := isExist[nums[i]]; !ok {
			isExist[nums[i]] = struct{}{}
		} else {
			continue
		}
		if nums[i] > minList[2] {
			minList[0] = minList[1]
			minList[1] = minList[2]
			minList[2] = nums[i]
			continue
		}
		if nums[i] > minList[1] {
			minList[0] = minList[1]
			minList[1] = nums[i]
			continue
		}
		if nums[i] > minList[0] {
			minList[0] = nums[i]
			continue
		}
	}
	if len(isExist) < 3 {
		for i := 0; i < len(minList); i++ {
			if r <= minList[i] {
				r = minList[i]
			}
		}
	} else {
		r = minList[0]
	}
	return r
}

// @lc code=end

