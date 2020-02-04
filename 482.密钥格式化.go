import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	S = strings.ToUpper(strings.Replace(S, "-", "", -1))
	r := ""
	firstLen := len(S) % K
	r = S[:firstLen]
	for i := firstLen; i+K <= len(S); i += K {
		if i != 0 {
			r = r + "-"
		}
		r += S[i : i+K]
	}
	return r
}

// @lc code=end

