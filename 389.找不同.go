/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	sList := map[rune]int{} // 出现次数
	for _, v := range s {
		sList[v] += 1
	}
	tList := map[rune]int{} // 出现次数
	for _, v := range t {
		tList[v] += 1
	}
	for k, v := range tList {
		if _, ok := sList[k]; ok {
			if sList[k] != v {
				return byte(k)
			}
		} else {
			return byte(k)
		}

	}
	return s[0]
}

// @lc code=end

