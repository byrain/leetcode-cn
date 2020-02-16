import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 */

// @lc code=start
func repeatedStringMatch(A string, B string) int {
	tempA := ""
	i := 1
	r := -1
	l := len(B)/len(A) + 2
	// fmt.Println(len(B), len(A))
	for i <= l {
		tempA = tempA + A
		if strings.Contains(tempA, B) {
			r = i
			break
		}
		i++
	}
	return r
}

// @lc code=end

