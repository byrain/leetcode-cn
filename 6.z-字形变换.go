/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	retSlice := make(map[int][]rune, numRows)
	i, flag := 0, -1
	for _, v := range s {
		retSlice[i] = append(retSlice[i], v)
		if i == numRows-1 || i == 0 {
			flag = 0 - flag
		}
		i += flag
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret += string(retSlice[i])
	}
	return ret
}

// @lc code=end

