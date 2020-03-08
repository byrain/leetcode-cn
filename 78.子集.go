/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start

var ret [][]int

func subsets(nums []int) [][]int {
	ret = [][]int{}

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
		path = append(path, nums[i])
		dfs(nums, path, k, i+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end

