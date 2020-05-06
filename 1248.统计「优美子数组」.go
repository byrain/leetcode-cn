import "fmt"

/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	dp := make([]int, 0)
	cnt, ret := 0, 0
	for i := 0; i < len(nums); i++ {
		cnt++
		if nums[i]%2 == 1 {
			dp = append(dp, cnt)
			cnt = 0
		}
		if len(dp) >= k {
			ret += dp[len(dp)-k]
		}
	}
	fmt.Println(dp)
	return ret
}

// @lc code=end

