import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := [][]int{}
	for i := range intervals {
		n := len(ret)
		if n == 0 || intervals[i][0] > ret[n-1][1] {
			ret = append(ret, intervals[i])
		} else {
			if ret[n-1][1] < intervals[i][1] {
				ret[n-1][1] = intervals[i][1]
			}
		}
	}
	return ret
}

// @lc code=end

