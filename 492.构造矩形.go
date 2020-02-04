import (
	"math"
)

/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	l := w
	for i := 0; i < area; i++ {
		if w == 0 || l > area {
			break
		}
		if w*l == area {
			return []int{l, w}
		} else if w*l < area {
			l++
		} else {
			w--
		}
	}
	return []int{l, w}
}

// @lc code=end

