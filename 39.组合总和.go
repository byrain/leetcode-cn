import "sort"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

var ret [][]int

func combinationSum(candidates []int, target int) [][]int {
	ret = [][]int{}
	sort.Ints(candidates)
	dfs(candidates, []int{}, target)
	return ret
}

func dfs(candidates []int, path []int, r int) {
	if r == 0 {
		pathTemp := make([]int, len(path))
		copy(pathTemp, path)
		ret = append(ret, pathTemp)
		return
	}
	if r < 0 {
		return
	}
	for _, v := range candidates {
		if len(path) > 0 && v < path[len(path)-1] {
			continue
		}

		path = append(path, v)
		dfs(candidates, path, r-v)
		path = path[:len(path)-1]
	}
}

// @lc code=end

