/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	dp[0] = make([]bool, len(s))

	ret := ""
	maxLen := 0
	j := 1
	for j < len(s) {
		dp[j] = make([]bool, len(s))
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}

			if j-i < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] {
				if (j - i + 1) > maxLen {
					maxLen = j - i + 1
					ret = s[i : j+1]
				}
			}
		}
		j++
	}
	if len(ret) == 0 {
		ret = string(s[0])
	}
	return ret
}

// @lc code=end

