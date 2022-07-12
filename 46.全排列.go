/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
var ret [][]int

func permute(nums []int) [][]int {
	ret = [][]int{}
	backtrace(nums, []int{})
	return ret
}

func backtrace(nums, path []int) {
	if len(nums) == 0 {
		newPath := make([]int, len(path))
		copy(newPath, path)
		ret = append(ret, newPath)
		return
	}

	for i := 0; i < len(nums); i++ {
		path = append(path, nums[i])
		newNums := []int{}
		newNums = append(newNums, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		backtrace(newNums, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end

