import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
func checkRecord(s string) bool {
	if strings.Count(s, "A") > 1 {
		return false
	}
	if strings.Count(s, "LLL") >= 1 {
		return false
	}
	return true
}

// @lc code=end

