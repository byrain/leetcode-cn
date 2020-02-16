/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if target < v {
			return v
		}
	}
	return letters[0]
}

// @lc code=end

