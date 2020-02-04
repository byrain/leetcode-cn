import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	if len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)
	p1 := 0
	p2 := 1
	r := map[int]struct{}{} // 保证结果不会重复
	for {
		if p1 == p2 { // 保证不会成对出现
			p2++
		}
		if p2 >= len(nums) {
			break
		}
		v := nums[p2] - nums[p1]
		switch {
		case v == k:
			r[nums[p1]] = struct{}{}
			p1++
			p2++
		case v < k:
			p2++
		case v > k:
			p1++
		}
	}
	return len(r)
}

// @lc code=end

