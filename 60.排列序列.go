/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
var ret [][]int

func getPermutation(n int, k int) string {
	ret = [][]int{}
	nums := []int{}
	for i := 1; i < n+1; i++ {
		nums = append(nums, i)
	}
	backtrace(nums, []int{}, k)
	var r strings.Builder
	for _, v := range ret[len(ret)-1] {
		fmt.Fprintf(&r, "%d", v)
	}

	return r.String()
}

func backtrace(nums, path []int, k int) {
	if len(nums) == 0 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		ret = append(ret, newPath)
		return
	}
	if len(ret) > k {
		return
	}
	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		newNums := []int{}
		newNums = append(newNums, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		backtrace(newNums, path, k)
		path = path[:len(path)-1]
	}
}

// @lc code=end

