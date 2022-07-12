import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start

var ret [][]int
var used map[int]bool

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return ret
	}
	ret = [][]int{}
	used = make(map[int]bool, len(nums))
	sort.Ints(nums)
	backtrace(nums, []int{})
	return ret
}

func backtrace(nums, path []int) {
	if len(path) == len(nums) {
		tempPath := make([]int, len(nums))
		copy(tempPath, path)
		ret = append(ret, tempPath)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrace(nums, path)
		used[i] = false
		path = path[:len(path)-1]
	}
}

// @lc code=end

