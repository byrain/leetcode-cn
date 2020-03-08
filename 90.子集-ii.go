import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start

var ret [][]int

func subsetsWithDup(nums []int) [][]int {
	ret = [][]int{}
	sort.Ints(nums) // 整理成已知问题
	for i := 0; i < len(nums)+1; i++ {
		dfs(nums, []int{}, i, 0)
	}
	return ret
}

func dfs(nums, path []int, k, p int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		ret = append(ret, temp)
		return
	}

	for i := p; i < len(nums); i++ {
		if i > p && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, path, k, i+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end

