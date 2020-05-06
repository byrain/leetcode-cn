/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	// sList := map[rune]int{} // 出现次数
	// for _, v := range s {
	// 	sList[v] += 1
	// }

	bs := []byte(s)
	bmap := make(map[byte]int, 52)
	for _, v := range bs {
		bmap[v] += 1
	}

	hasOdd := false
	r := 0
	for _, v := range bmap {
		if v%2 == 0 {
			r += v
		} else {
			r += v - 1
			hasOdd = true
		}
	}
	if hasOdd {
		r += 1
	}
	return r
}

// @lc code=end

