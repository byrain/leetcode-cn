import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	ret := 0
	if len(nums) < 3 {
		return ret
	}
	distance := math.MaxInt32
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			absDistance := int(math.Abs(float64(target - sum)))
			if absDistance <= distance {
				distance = absDistance
				ret = sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}

	return ret
}

// @lc code=end

