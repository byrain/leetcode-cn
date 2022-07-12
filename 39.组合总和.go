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
		if len(path) > 0 && candidates[i] < path[len(path)-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrace(candidates, path, target-candidates[i])
		path = path[:len(path)-1]
	}
}

// @lc code=end
