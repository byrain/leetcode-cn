/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	magList := map[rune]int{}
	for _, v := range magazine {
		magList[v] += 1
	}

	for _, v := range ransomNote {
		if _, ok := magList[v]; !ok {
			return false
		} else {
			if magList[v] == 0 {
				return false
			} else {
				magList[v] -= 1
			}
		}
	}
	return true
}

// @lc code=end

