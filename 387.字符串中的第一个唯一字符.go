/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	magList := map[rune]int{} // 出现次数
	for _, v := range s {
		magList[v] += 1
	}
	for k, v := range s {
		if magList[v] == 1 {
			return k
		}
	}
	return -1
}

// @lc code=end

