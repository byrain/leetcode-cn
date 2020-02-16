/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(str string) string {
	temp := []rune{}
	for _, v := range str {
		if v >= 65 && v <= 90 {
			temp = append(temp, v+32)
		} else {
			temp = append(temp, v)
		}
	}
	return string(temp)
}

// @lc code=end

