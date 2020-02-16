/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	degreeNums := []int{}
	degree := map[int]int{}
	left, right := map[int]int{}, map[int]int{}
	ret := 50000
	maxDegree := 0
	for k, v := range nums {
		if _, ok := left[v]; !ok {
			left[v] = k
		}
		right[v] = k
		degree[v]++
		if degree[v] > maxDegree {
			maxDegree = degree[v]
		}
	}

	for k, v := range degree {
		if v == maxDegree {
			degreeNums = append(degreeNums, k)
		}
	}

	for _, v := range degreeNums {
		r := right[v] - left[v] + 1
		if r < ret {
			ret = r
		}
	}
	return ret
}

// @lc code=end

