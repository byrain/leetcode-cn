import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
var ret [][]int

func combinationSum2(candidates []int, target int) [][]int {
	ret = [][]int{}
	sort.Ints(candidates)
	dfs(candidates, []int{}, target, 0)
	return ret
}

func dfs(candidates []int, path []int, r int, p int) {
	if r == 0 {
		pathTemp := make([]int, len(path))
		copy(pathTemp, path)
		ret = append(ret, pathTemp)
		return
	}
	if r < 0 {
		return
	}
	for k, v := range candidates {
		if len(path) > 0 && k < p {
			continue
		}

		if k > p && candidates[k] == candidates[k-1] {
			continue
		}

		path = append(path, v)
		dfs(candidates, path, r-v, k+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end

