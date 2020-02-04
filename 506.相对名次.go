import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
func findRelativeRanks(nums []int) []string {
	r := []string{}
	m := map[int]int{}
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	for i := len(temp) - 1; i >= 0; i-- {
		m[temp[i]] = len(temp) - i
	}

	for i := 0; i < len(nums); i++ {
		switch m[nums[i]] {
		case 1:
			r = append(r, "Gold Medal")
		case 2:
			r = append(r, "Silver Medal")
		case 3:
			r = append(r, "Bronze Medal")
		default:
			r = append(r, strconv.Itoa(m[nums[i]]))
		}
	}
	return r
}

// @lc code=end

