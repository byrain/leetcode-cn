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
var used map[int]bool

func combinationSum2(candidates []int, target int) [][]int {
	ret = [][]int{}
	used = map[int]bool{}
	sort.Ints(candidates)
	backtrace(candidates, []int{}, target)
	return ret
}

func backtrace(candidates []int, path []int, target int) {
	if target == 0 {
		pathTemp := make([]int, len(path))
		copy(pathTemp, path)
		ret = append(ret, pathTemp)
		return
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		if used[i] {
			continue
		}
		if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
			continue
		}
		if len(path) > 0 && candidates[i] < path[len(path)-1] {
			continue
		}
		path = append(path, candidates[i])
		used[i] = true
		backtrace(candidates, path, target-candidates[i])
		used[i] = false
		path = path[:len(path)-1]
	}
}

// @lc code=end

