/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start

var ret [][]int

func combine(n int, k int) [][]int {
	ret = [][]int{}
	nums := []int{}
	for i := 1; i < n+1; i++ {
		nums = append(nums, i)
	}

	if k > n {
		ret = append(ret, nums)
		return ret
	}

	dfs(nums, []int{}, k, 0)
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

