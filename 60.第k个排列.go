import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	nums := make([]int, n)
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1] * i
		nums[i-1] = i
	}
	k--
	var b strings.Builder
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]
		fmt.Fprintf(&b, "%d", nums[a])
		nums = append(nums[:a], nums[a+1:]...)
	}
	return b.String()
}

// @lc code=end

