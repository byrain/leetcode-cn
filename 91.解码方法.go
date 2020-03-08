/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	size := len(s)
	if size == 0 || s == "0" || int(s[0])-48 == 0 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < size; i++ {
		t := int(s[i]) - 48
		if t != 0 {
			dp[i] = dp[i-1]
		}
		if i >= 1 {
			t = (int(s[i-1])-48)*10 + (int(s[i]) - 48)
			if t >= 10 && t <= 26 {
				if i == 1 {
					dp[i] += 1
				} else {
					dp[i] += dp[i-2]
				}
			}
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

