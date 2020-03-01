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
	dfs(nums, []int{})
	return ret
}

func dfs(nums, path []int) {
	if len(nums) == len(path) {
		newPath := make([]int, len(path))
		copy(newPath, path)
		ret = append(ret, newPath)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		path = append(path, nums[i])
		used[i] = true
		dfs(nums, path)
		used[i] = false
		path = path[:len(path)-1]
	}
}

// @lc code=end

